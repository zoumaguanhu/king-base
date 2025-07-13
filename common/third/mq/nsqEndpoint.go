package mq

import (
	"context"
	"github.com/golang/snappy"
	"github.com/nsqio/go-nsq"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"king.com/king/base/common/constants"
	"king.com/king/base/common/strs"
	"math/rand"
	"time"
)

type NsqTopicConfig struct {
	NsqdAddrs           []string `json:"NsqdAddrs"`    // 支持多个nsqd
	LookupdAddrs        []string `json:"LookupdAddrs"` // 支持多个lookupd
	Name                string   `json:"Name"`
	Channels            []string `json:"Channels"`
	MsgTimeout          int      `json:"MsgTimeout,default=10"`         // 秒
	MaxInFlight         int      `json:"MaxInFlight,default=25"`        // 并发数
	DefaultRequeueDelay int      `json:"DefaultRequeueDelay,default=3"` // 秒
	MaxAttempts         int      `json:"MaxAttempts,default=3"`         // 最大重试次数
}
type NsqConfig struct {
	Topics []NsqTopicConfig `json:"Topics"`
}
type NsqProducer struct {
	producer *nsq.Producer
}

func createNsqConfig(c *NsqTopicConfig) *nsq.Config {
	config := nsq.NewConfig()

	// 消息处理超时
	config.MsgTimeout = time.Duration(c.MsgTimeout) * time.Second

	// 最大飞行中消息数(控制并发)
	config.MaxInFlight = c.MaxInFlight

	// 重试策略
	config.DefaultRequeueDelay = time.Duration(c.DefaultRequeueDelay) * time.Second
	config.MaxAttempts = uint16(c.MaxAttempts)

	return config
}

type NsqProducerWrapper struct {
	nsp *map[string]*[]*nsq.Producer
}

func NewNsqProducerWrapper(c *NsqConfig) (*NsqProducerWrapper, bool) {
	pmp, err := NewNsqProducer(c)
	if err != nil {
		return nil, false
	}
	return &NsqProducerWrapper{nsp: pmp}, true
}

func (w *NsqProducerWrapper) Publish(topic string, body []byte) error {
	if ns, ok := (*w.nsp)[topic]; ok {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(*ns))
		if e := w.randomPublish((*ns)[randomIndex], topic, body); e != nil {
			for _, producer := range *ns {
				if e1 := w.randomPublish(producer, topic, body); e1 != nil {
					logx.Errorf("NSQ Publish topic:%v,err1:%v", topic, e)
					continue
				}
				logx.Infof("NSQ Publish topic:%v success", topic)
				return nil
			}
			logx.Errorf("NSQ Publish topic:%v,err:%v", topic, e)
			return e
		}
		logx.Infof("NSQ Publish topic:%v success", topic)

	}
	return nil
}
func (w *NsqProducerWrapper) randomPublish(nsp *nsq.Producer, topic string, body []byte) error {
	if e := nsp.Publish(topic, body); e != nil {
		logx.Errorf("nsq producer err:%v", e)
		return e
	}
	return nil
}
func NewNsqProducer(c *NsqConfig) (*map[string]*[]*nsq.Producer, error) {
	config := nsq.NewConfig()
	nmp := make(map[string]*[]*nsq.Producer, len(c.Topics))

	// 手动连接所有nsqd节点
	for _, topic := range c.Topics {
		nps := &[]*nsq.Producer{}
		for _, addr := range topic.NsqdAddrs {
			producer, err := nsq.NewProducer(addr, config)
			if err != nil {
				continue
			}
			if err := producer.Ping(); err != nil {
				logx.Errorf("connection nsqd fail: %s, error: %v", addr, err)
				continue
			}
			logx.Infof("Successful connection nsqd: %s", addr)
			*nps = append(*nps, producer)
		}
		if len(*nps) <= 0 {
			continue
		}
		nmp[topic.Name] = nps
	}
	return &nmp, nil
}

// 生产者健康检查
func (p *NsqProducer) HealthCheck() bool {
	return p.producer.Ping() == nil
}
func (p *NsqProducer) Publish(topic string, body []byte) error {
	return p.producer.Publish(topic, body)
}
func (p *NsqProducer) PublishCompressed(topic string, body []byte) error {
	compressed := snappy.Encode(nil, body)
	return p.producer.Publish(topic, compressed)
}

func (p *NsqProducer) Stop() {
	p.producer.Stop()
}

type NsqHandler struct {
	ctx    context.Context
	handle func(ctx context.Context, message []byte) error
}

func NewNsqHandler(handle func(ctx context.Context, message []byte) error) *NsqHandler {

	return &NsqHandler{
		handle: handle,
	}
}

func (h *NsqHandler) HandleMessage(msg *nsq.Message) error {
	defer func() {
		if e := recover(); e != nil {
			logc.Errorf(h.ctx, "HandleMessage err:%v", e)
		}
	}()
	logx.WithContext(h.ctx).Infof("Received the NSQ message: %s", string(msg.Body))
	s := string(msg.Body)
	ms := &MsgStruct{}
	strs.StrToObj(&s, ms)
	ctx := logx.WithFields(context.Background(), logx.Field(constants.TRACE_ID, ms.Header.MsgId))
	h.ctx = ctx
	msg.DisableAutoResponse()
	if err := h.handle(h.ctx, msg.Body); err != nil {
		logx.WithContext(h.ctx).Errorf("Message processing failed: %v", err)
		return err // 返回错误会触发NSQ的自动重试
	}

	return nil
}

type NsqConsumerWrapper struct {
	ncs    *[]*NsqConsumer
	config *NsqConfig
}

func NewNsqConsumerWrapper(c *NsqConfig) (*NsqConsumerWrapper, bool) {
	nmp := &[]*NsqConsumer{}
	return &NsqConsumerWrapper{ncs: nmp, config: c}, true
}

type NsqConsumer struct {
	consumer *nsq.Consumer
	handler  *NsqHandler
	channel  string
	topic    string
}

func (c *NsqConsumerWrapper) AddConsumer(nc *NsqConsumer) {
	*c.ncs = append(*c.ncs, nc)
}
func (c *NsqConsumerWrapper) Start(c1 *NsqConfig) {

}

func (c *NsqConsumerWrapper) NewNsqConsumer(chl string, handler func(ctx context.Context, message []byte) error) (*NsqConsumer, error) {
	for _, topic := range c.config.Topics {
		for _, channel := range topic.Channels {
			if channel != chl {
				continue
			}
			consumer, err := nsq.NewConsumer(topic.Name, channel, createNsqConfig(&topic))
			if err != nil {
				logx.Errorf("NewNsqConsumer topic:%v,channel:%v, err:%v", topic.Name, channel, err)
				return nil, err
			}

			consumer.AddHandler(NewNsqHandler(handler))
			for _, addr := range topic.LookupdAddrs {
				// 连接NSQLookupd发现服务
				if err := consumer.ConnectToNSQLookupd(addr); err != nil {
					logx.Errorf("NewNsqConsumer ConnectToNSQLookupd add:%v,err:%v", addr, err)
					return nil, err
				}
			}

			nc := &NsqConsumer{channel: channel, topic: topic.Name, consumer: consumer}
			return nc, nil
		}
	}
	return nil, nil
}

func (c *NsqConsumer) Start(cg *NsqConfig) {
	// 已在ConnectToNSQLookupd时启动

}

func (c *NsqConsumer) Stop() {

}
