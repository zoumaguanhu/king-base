package mq

type MsgHeader struct {
	MsgId        string
	ExchangeName string
	RoutingKey   string
	PublishTime  string
}
type MsgBody struct {
	TraceId     string
	CommandType string
	MsgContent  string
}
type MsgStruct struct {
	Header *MsgHeader
	Body   *MsgBody
}
