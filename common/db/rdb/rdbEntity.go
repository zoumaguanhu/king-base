package rdb

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
