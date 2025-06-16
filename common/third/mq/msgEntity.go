package mq

import (
	"king.com/king/base/common/third/mail"
)

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
	PublishId   int64  `json:"publishId"`
	IsBatch     bool   `json:"isBatch"`
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
	MailConfig    *mail.MailConf `json:"mailConfig"`
	MailContent   *mail.Email    `json:"mailContent"`
	Account       string         `json:"account"`
	IsHtml        bool           `json:"isHtml"`
	MailType      int64          `json:"mailType"`
	VirtId        int64          `json:"virtId"`
	PId           int64          `json:"pId"`    //券ID
	AcType        string         `json:"acType"` //活动类型
	AcCode        string         `json:"acCode"` //活动code
	AcId          int64          `json:"acId"`
	IssueUserId   int64          `json:"issueUserId"`
	TraceId       string         `json:"traceId"`
	IsSendMail    bool           `json:"isSendMail"`
	IssueDateTime string         `json:"issueDateTime"` //发放时间
	Amount        int64          `json:"amount"`        //消费金额字段
	IssueType     int64          `json:"issueType"`
	InviteCode    string         `json:"inviteCode"`
}
type CartInfo struct {
	Id         int64 `json:"id"`     // id
	PId        int64 `json:"pId"`    // p_id
	PCount     int64 `json:"pCount"` // p_count
	Status     int64 `json:"status"` // status
	SkuId      int64 `json:"skuId"`
	SkuChildId int64 `json:"skuChildId"`
	MId        int64 `json:"MId"`
}
type UserAction struct {
	MId    int64  `json:"mId"`
	VirtId int64  `json:"virtId"`
	Exp    string `json:"exp"`
}
