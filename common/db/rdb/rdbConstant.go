package rdb

import "time"

const (
	CLIENT_STAT_PV = "stat_pv"
)
const (
	SITE         = "v_site_info"
	HOST         = "host"
	CLIENT       = "client"
	CART         = "cart"
	STAT         = "stat"
	PRODUCT      = "product"
	PRODUCT_HASH = "product_hash"
	PRODUCT_SET  = "product_set"
	BANNER       = "banner"
	FILE_LIMIT   = "file_limit"
	MAIL_CODE    = "mail_code"
	ADMIN_CLIENT = "admin_client"
	ADMIN_CODE   = "admin_code"
)
const (
	PRODUCT_KEY        = "site:vsite:%v:product:page:product_hash"
	PRODUCT_SKU_KEY    = "site:vsite:%v:product:detail:%v:product_skus"
	PRODUCT_IMG_KEY    = "site:vsite:%v:product:detail:%v:product_imgs"
	PRODUCT_REVIEW_KEY = "site:vsite:%v:product:detail:%v:product_review"
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
	DANGER_ADMIN_LIMIT        = 8
	ADMIN_TRY_COUNT_LIMIT int = 5
)

// product缓存bz关键字
const (
	REVIEW_PAGE = "review_page"
)

// user缓存bz关键字
const (
	ADMIN_CODE_EXP          int64 = 300
	ADMIN_LOGIN_REPEAT_EXP  int64 = 1
	ADMIN_LOGIN_RETRY_COUNT int64 = 2
	ADMIN_CODE_RETRY_EXP    int64 = 50
	ADMIN_CODE_RETRY_COUNT  int64 = 2
)
const (
	USER_INFO            string = "user_info"
	USER_TEMP            string = "user_temp"
	USER_TEMP_EXP        int64  = 1
	USER_TEMP_FLG        string = "1"
	USER_TRY_COUNT       string = "try_count"
	USER_TRY_COUNT_LIMIT int    = 5
)
