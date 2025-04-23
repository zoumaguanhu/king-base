package rdb

type RdbIBase interface {
	BuildSiteKey(virtId int64, bz string) string
	BuildUserKey(virtId int64, userId int64, bz string) string
	BuildCache(k string, v interface{})
}
