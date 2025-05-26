package mq

import "king.com/king/base/common/third/mail"

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
	VirtId      int64  `json:"virtId"`
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
type MsgEmail struct {
	MailConfig  *mail.MailConf `json:"mailConfig"`
	MailContent *mail.Email    `json:"mailContent"`
	MailType    int64          `json:"mailType"`
	VirtId      int64          `json:"virtId"`
	TraceId     string         `json:"traceId"`
}
type MsgCoupon struct {
	MailConfig  *mail.MailConf `json:"mailConfig"`
	MailContent *mail.Email    `json:"mailContent"`
	Accounts    []string       `json:"accounts"`
	IsHtml      bool           `json:"isHtml"`
	MailType    int64          `json:"mailType"`
	VirtId      int64          `json:"virtId"`
	PId         int64          `json:"pId"` //券ID
	ACode       string         `json:"aId"` //活动ID
	IssueUserId int64          `json:"issueUserId"`
	TraceId     string         `json:"traceId"`
	IsSendMail  bool           `json:"isSendMail"`
}
