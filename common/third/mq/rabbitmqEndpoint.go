package mq

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zeromicro/go-zero/core/logx"
	"king.com/king/base/common/constants"
	"log"
	"sync"
	"time"
)

type ConsumerFun func()

type RabbitMQ struct {
	URL       string
	VHost     string
	Exchanges []ExchangeConfig
	Queues    []QueueConfig
}
type ExchangeConfig struct {
	Name    string
	Type    string
	BootDLX bool //是否开启死信交换机
	Durable bool
	Desc    *string //描述信息
}

type QueueConfig struct {
	Name                 string
	Exchange             string
	RoutingKey           string
	Durable              bool
	BootDLX              bool  //是否开启死信队列
	QueueConsumerTimeout int64 //队列超时时间，单位秒
	MaxRetryCount        int64 //最大重试次数，达到次数进入死信队列
}
type RabbitMQConsumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	msgs    <-chan amqp.Delivery
	mu      sync.Mutex
	done    chan struct{}
}

type RabbitMQEndpoint struct {
	conn         *amqp.Connection
	channel      *amqp.Channel
	queues       map[string]amqp.Queue
	config       *RabbitMQ
	exchangeMap  *map[string]*ExchangeConfig
	queueMap     *map[string]*QueueConfig
	consumerList *[]ConsumerFun
	done         chan struct{}
}

func NewRabbitMQEndpoint(mq *RabbitMQ) (*RabbitMQEndpoint, error) {
	config := amqp.Config{Vhost: mq.VHost, Locale: "en_US", Heartbeat: 20 * time.Second}
	conn, err := amqp.DialConfig(mq.URL, config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	queues := make(map[string]amqp.Queue)
	exchangeMap := make(map[string]*ExchangeConfig)
	queueMap := make(map[string]*QueueConfig)

	// 声明交换机
	for _, exchangeConf := range mq.Exchanges {
		if err := exchangeBind(ch, conn, exchangeConf.Name, exchangeConf.Type, exchangeConf.Durable); err != nil {
			return nil, err
		}
		if exchangeConf.BootDLX {
			if err := exchangeBind(ch, conn, constants.DLX_PREFIX+exchangeConf.Name, constants.DIRECT, exchangeConf.Durable); err != nil {
				return nil, err
			}
		}
		exchangeMap[exchangeConf.Name] = &exchangeConf
	}

	// 声明队列并绑定到交换机
	for _, queueConf := range mq.Queues {
		args := amqp.Table{}
		args["x-queue-consumer-timeout"] = int32(queueConf.QueueConsumerTimeout * time.Second.Milliseconds())
		args["max-retry-count"] = int32(queueConf.MaxRetryCount)
		args["boot-dlx"] = queueConf.BootDLX
		//配置中的队列绑定
		q, err := queueBind(ch, conn, queueConf.Name, queueConf.Exchange, queueConf.RoutingKey, queueConf.Durable, queueConf.BootDLX, args)
		if err != nil {
			return nil, err
		}
		queues[queueConf.Name] = *q
		//绑定死信队列
		if queueConf.BootDLX {
			_, err := queueBind(ch, conn, constants.DLX_PREFIX+queueConf.Name, constants.DLX_PREFIX+queueConf.Exchange, constants.DLX_PREFIX+queueConf.RoutingKey, queueConf.Durable, false, nil)
			if err != nil {
				return nil, err
			}
		}
		queueMap[queueConf.Name] = &queueConf
	}

	return &RabbitMQEndpoint{
		conn:         conn,
		channel:      ch,
		queues:       queues,
		config:       mq,
		exchangeMap:  &exchangeMap,
		queueMap:     &queueMap,
		consumerList: &[]ConsumerFun{},
	}, nil
}
func (s RabbitMQEndpoint) StartConsumer() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logx.Errorf("StartConsumer err:%v", r)
			}
		}()
		for _, f := range *s.consumerList {
			//启动消费者
			f()
		}
	}()
}
func (s *RabbitMQEndpoint) ConsumeMessages(queueName string, handler func(d amqp.Delivery) error) error {
	if !s.IsConnected() {
		logx.Errorf("rabbitmq lost connect")
		return nil
	}
	queue, ok := s.queues[queueName]
	if !ok {
		return nil
	}
	chs, err := s.channel.Consume(
		queue.Name,
		"",
		false, // 禁用自动确认，开启手动 ACK
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logx.Errorf("failed to register a consumer: %w", err)
		return err
	}
	sem := make(chan struct{}, 10)

	for d := range chs {
		sem <- struct{}{}
		logx.Infof("get goroutine from sem limit :%v", len(sem))
		go func(m *amqp.Delivery) {
			defer func() {
				if r := recover(); r != nil {
					logx.Errorf("panic err: %v\n", r)
				}
				<-sem
			}()
			if e := handler(*m); e != nil {
				s.maxRetryCount(queueName, m)
			}

		}(&d)
	}

	return nil
}
func (s *RabbitMQEndpoint) RegisterConsumer(c ConsumerFun) {
	*s.consumerList = append(*s.consumerList, c)
}
func (s *RabbitMQEndpoint) RegisterConsumers(c ...ConsumerFun) {
	*s.consumerList = append(*s.consumerList, c...)
}
func exchangeBind(ch *amqp.Channel, conn *amqp.Connection, exchange, exType string, durable bool) error {
	err := ch.ExchangeDeclare(
		exchange,
		exType,
		durable,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		logx.Errorf("failed to declare exchange %s: %w", exchange, err)
		return err
	}
	return nil
}
func queueBind(ch *amqp.Channel, conn *amqp.Connection, queueName, exchange, routingKey string, durable, bootDLX bool, args amqp.Table) (*amqp.Queue, error) {
	if bootDLX {
		args["x-dead-letter-exchange"] = "dlx_" + exchange
		args["x-dead-letter-routing-key"] = "dlx_" + routingKey
	}
	q, err := ch.QueueDeclare(
		queueName,
		durable,
		false,
		false,
		false,
		args,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		logx.Errorf("failed to declare queue %s: %w", queueName, err)
		return nil, err
	}

	if err = ch.QueueBind(
		q.Name,
		routingKey,
		exchange,
		false,
		nil,
	); err != nil {
		ch.Close()
		conn.Close()
		logx.Errorf("failed to bind queue %s to exchange %s: %w", queueName, exchange, err)
		return nil, err
	}
	return &q, nil
}

func (s *RabbitMQEndpoint) Close() {
	s.channel.Close()
	s.conn.Close()
}
func (s *RabbitMQEndpoint) SendMessage(msg *MsgStruct) error {
	if !s.IsConnected() {
		logx.Errorf("rabbitmq lost connect")
		return nil
	}
	logx.Infof("send msg:%v", *msg)
	header := msg.Header
	header.PublishTime = time.Now().String()
	body, err1 := json.Marshal(msg.Body)
	if err1 != nil {
		logx.Errorf("failed msg body to json %s: %w", header.ExchangeName, err1)
		return err1
	}
	headers := amqp.Table{}
	if exchange, ok := (*s.exchangeMap)[header.ExchangeName]; ok {
		if exchange.BootDLX {
			headers["retry_count"] = 0
		}
	}

	if err := s.channel.Publish(
		header.ExchangeName,
		header.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "text/plain",
			DeliveryMode: amqp.Persistent,
			Body:         body,
			Headers:      headers,
		}); err != nil {
		logx.Errorf("failed to publish a message to exchange %s: %w", header.ExchangeName, err)
	}
	return nil
}

func (s *RabbitMQEndpoint) maxRetryCount(queueName string, m *amqp.Delivery) bool {
	retryCount, ok := m.Headers["retry_count"].(int32)
	if !ok {
		return false
	}
	if queue, ok := (*s.queueMap)[queueName]; ok {
		retryCount++
		if retryCount >= int32(queue.MaxRetryCount) {
			// 达到最大重试次数，拒绝消息且不重新入队，消息将进入死信队列
			if err := m.Reject(false); err != nil {
				logx.Infof("Failed to reject message: %s", err)
			}
			logx.Infof("Message reached max retry count, sent to DLX")
			return true
		} else {
			m.Headers["retry_count"] = retryCount
			err := s.channel.Publish(
				m.Exchange,
				m.RoutingKey,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        m.Body,
					Headers:     m.Headers,
				},
			)
			if err != nil {
				logx.Errorf("Failed to republish message: %s", err)
			}
			// 确认原始消息，避免重复处理
			err = m.Ack(false)
			if err != nil {
				logx.Errorf("Failed to ack message: %s", err)
			}
		}
	}

	return false
}

// IsConnected 验证连接是否存在
func (s *RabbitMQEndpoint) IsConnected() bool {
	return s.conn != nil && !s.conn.IsClosed() && s.channel != nil && !s.channel.IsClosed()
}
func (s *RabbitMQEndpoint) reConnected() error {
	logx.Infof("rabbitmq reconnected")
	var err error
	s.Close()

	endpoint, err := NewRabbitMQEndpoint(s.config)
	if err != nil {
		logx.Errorf("reconnected err:%v", err)
		return err
	}
	s.conn = endpoint.conn
	s.channel = endpoint.channel
	s.queues = endpoint.queues
	return nil
}

// monitorConnection 监控连接状态
func (s *RabbitMQEndpoint) MonitorConnection() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logx.Errorf("rabbit monitorConnection exit:%v", r)
			}
		}()
		notifyClose := make(chan *amqp.Error)
		s.conn.NotifyClose(notifyClose)

		for {
			time.Sleep(5)
			select {
			case <-s.done:
				return
			case err := <-notifyClose:
				if err != nil {
					log.Printf("RabbitMQ connection closed: %v. Reconnecting...", err)
					for {
						endpoint, err := NewRabbitMQEndpoint(s.config)
						if err == nil {
							log.Println("Reconnected to RabbitMQ successfully.")
							s.conn = endpoint.conn
							s.channel = endpoint.channel
							s.queues = endpoint.queues
							notifyClose = make(chan *amqp.Error)
							s.conn.NotifyClose(notifyClose)
							//重新开始消费
							s.StartConsumer()
							break
						}
						log.Printf("Failed to reconnect to RabbitMQ: %v. Retrying in %s...", err, 5)
					}
				}
			}
		}
	}()

}
