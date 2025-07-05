package rdb

import "time"

// 用户缓存信息结构体

type UserInfo struct {
	Id        int64  `json:"id"`
	EMail     string `json:"EMail"`
	Status    int64  `json:"status"`
	MTag      string `json:"MTag"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Source    string `json:"source"`
	VirtId    int64  `json:"virtId"`
	BuildTime string `json:"buildTime"`
}

type SiteInfo struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	SiteUrl      string `json:"siteUrl"`
	SiteLogo     string `json:"siteLogo"`
	SiteMeta     string `json:"siteMeta"`
	Status       int64  `json:"status"`
	Zone         int64  `json:"zone"`
	SiteTemplate string `json:"siteTemplate"`
	ExpTime      string `json:"expTime"`
	Lang         string `json:"lang"`
	BuildTime    string `json:"buildTime"`
}
type SiteStat struct {
	Count int64 `json:"count"`
}
type SiteMailLimit struct {
	Count           int64 `json:"count"`
	MaxCountPerDay  int64 `json:"maxCountPerDay"`
	MaxCountPerFile int64 `json:"maxCountPerFile"`
}

type BannerInfo struct {
	Id         int64  `json:"id"`
	BannerUrl  string `json:"bannerUrl"`
	RouteToUrl string `json:"routeToUrl"`
	BSort      int64  `json:"BSort"`
}
type MailCode struct {
	Code      string `json:"code"`
	BuildTime int64  `json:"buildTime"`
}
type ContentSimpleTp struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	PublishTime string `json:"publishTime"`
	VirtId      int64  `json:"virtId"`
}
type ContentDetailTp struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	PublishTime string `json:"publishTime"`
	VirtId      int64  `json:"virtId"`
}
type ContentTpLimit struct {
	ContentType  int64 `json:"contentType"`
	CreateLimit  int64 `json:"createLimit"`
	PublishLimit int64 `json:"publishLimit"`
	Status       int64 `json:"status"`
}
type OrderContent struct {
	OrderNo      string    `json:"orderNo"`
	Status       int64     `json:"status"`
	OrderType    int64     `json:"orderType"`
	TotalPrice   string    `json:"totalPrice"`
	RealPrice    string    `json:"realPrice"`
	OrderTimeOut time.Time `json:"orderTimeout"`
	MId          int64     `json:"mId"`
	XCode        string    `json:"xCode"`
	Platform     string    `json:"platform"`
	VirtId       int64     `json:"virtId"`
	CartStep     int64     `json:"cartStep"`  //阶段参数
	OrderStep    int64     `json:"orderStep"` //阶段参数
}
