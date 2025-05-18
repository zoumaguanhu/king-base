package rdb

import "time"

const (
	CLIENT_STAT_PV = "stat_pv"
)
const (
	SITE    = "v_site_info"
	HOST    = "host"
	CLIENT  = "client"
	STAT    = "stat"
	PRODUCT = "product"
)
const (
	STAT_PV   = "pv"
	STAT_UV   = "uv"
	SYNC_TIME = "sync_time"
)
const (
	STAT_TYPE_1 int64 = 1
	STAT_TYPE_2 int64 = 2
	STAT_TYPE_3 int64 = 3
	STAT_TYPE_4 int64 = 4
)
const (
	STAT_EXP_TIME             = time.Duration(604800 * time.Second) // 7 * 24 * 60 * 60
	DANGER_CLIENT_LIMIT       = 6
	DANGER_ADMIN_LIMIT        = 6
	ADMIN_TRY_COUNT_LIMIT int = 5
)

// product缓存bz关键字
const (
	REVIEW_PAGE = "review_page"
)

// user缓存bz关键字
const (
	USER_INFO            string = "user_info"
	USER_TEMP            string = "user_temp"
	USER_TEMP_EXP        int64  = 1
	USER_TEMP_FLG        string = "1"
	USER_TRY_COUNT       string = "try_count"
	USER_TRY_COUNT_LIMIT int    = 5
)
