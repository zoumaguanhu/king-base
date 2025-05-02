package rdb

import "time"

const (
	CLIENT_STAT_PV = "stat_pv"
)
const (
	SITE = "v_site_info"
)
const (
	STAT_PV = "pv"
	STAT_UV = "uv"
)
const (
	STAT_EXP_TIME = time.Duration(1296000 * time.Second) // 15 * 24 * 60 * 60
)

// product缓存bz关键字
const (
	REVIEW_PAGE = "review_page"
)

// user缓存bz关键字
const (
	USER_INFO = "user_info"
)
