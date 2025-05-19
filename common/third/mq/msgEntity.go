package mq

type MsgHeader struct {
	MsgId        string `json:"msgId"`
	ExchangeName string `json:"exchangeName"`
	RoutingKey   string `json:"routingKey"`
	PublishTime  string `json:"publishTime"`
}
type MsgBody struct {
	TraceId     string `json:"traceId"`
	CommandType string `json:"commandType"`
	Bz          string `json:"bz"`
	MsgContent  string `json:"msgContent"`
}
type MsgStruct struct {
	Header *MsgHeader
	Body   *MsgBody
}
type MsgProduct struct {
	Id      int64 `json:"id"`
	VirtId  int64 `json:"virtId"`
	Updater int64 `json:"updater"`
}
type MsgSite struct {
	Id      int64 `json:"id"`
	Updater int64 `json:"updater"`
}
type MsgBanner struct {
	Id      int64 `json:"id"`
	VirtId  int64 `json:"virtId"`
	Updater int64 `json:"updater"`
}
