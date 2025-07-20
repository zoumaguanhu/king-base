package rdb

import "time"

const (
	CLIENT_STAT_PV = "stat_pv"
	CLIENT_XCODE   = "x_code"
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
	MAIL_LIMIT   = "mail_limit"
	MAIL_CODE    = "mail_code"
	ADMIN_CLIENT = "admin_client"
	ADMIN_CODE   = "admin_code"
	SITE_BLOG    = "blog"
)
const (
	PRODUCT_KEY           = "site:vsite:%v:product:page:product_hash"
	PRODUCT_SKU_KEY       = "site:vsite:%v:product:detail:%v:product_skus"
	PRODUCT_IMG_KEY       = "site:vsite:%v:product:detail:%v:product_imgs"
	PRODUCT_REVIEW_KEY    = "site:vsite:%v:product:detail:%v:product_review"
	PRODUCT_STORE_KEY     = "site:vsite:%v:product:store:%v:%v:product_store"
	CARK_HASH_KEY         = "site:vsite:%v:member:%v:cart:cart_hash"
	ADDRESS_PREFIX_KEY    = "site:vsite:%v:member:%v:address:address"
	ADDRESS_HASH_KEY      = "site:vsite:%v:member:%v:address:address_hash"
	PRODUCT_STORE_OPTIONE = "%v_%v"
	BLOG_PAGE_HASH_KEY    = "site:vsite:%v:blog:page:blog_hash"
	BLOG_PAGE_SET_KEY     = "site:vsite:%v:blog:page:blog_set"
	BLOG_PAGE_PREFIX_KEY  = "site:vsite:%v:blog:page:blog"
	BLOG_DETAIL_KEY       = "site:vsite:%v:blog:detail:blog_hash"
	ACCOUNT_CODE_KEY      = "site:vsite:%v:code:account:%v"
)
const (
	STAT_FIELD_XCODE   = "platform:%v:bz:%v:xCode:%v"
	STAT_FIELDE        = "platform:%v:bz:%v"
	SITE_OPTIONS       = "site:vsite:%v"
	SITE_OPTIONS_FIELD = "option_%v"
	HOST_STAT_SUFFIX   = "bz:%v:date:%v"
)
const (
	ORDER_BASE_INFO  = "base_info"
	ORDER_BASE_ITEMS = "items_info"

	ORDER_TEMP_KEY        = "site:vsite:%v:member:%v:order:%v"
	ORDER_PAGE_KEY_PREFIX = "site:vsite:%v:member:%v:order:page:order"
	ORDER_DETAIL          = "site:vsite:%v:member:%v:order:detail:%v"
	CART_TEMP_KEY         = "site:vsite:%v:member:%v:cart:%v"
)
const (
	COUPONE_PAGE_KEY_PREFIX = "site:vsite:%v:member:%v:coupon:page:coupon"
	COUPONE_PAGE_KEY        = "site:vsite:%v:member:%v:coupon:page:coupon_hash"
)
const (
	COUNTRY_LIST_HASH = "site:vsite:common:country:country_hash"
	COUNTRY_LIST_INFO = "site:vsite:common:country:country_info"
	PAY_METHOD_HASH   = "site:vsite:common:pay_method:pay_hash"
	PAY_METHOD_INFO   = "site:vsite:common:pay_method:pay_info"
	LOCATION_API_INFO = "site:vsite:common:location_api:location_info"
)
const (
	STAT_PV      = "pv"
	STAT_UV      = "uv"
	SYNC_TIME    = "sync_time"
	NATURAL_FLOW = "0000"
)
const (
	STAT_TYPE_1 int64 = 1 //uv
	STAT_TYPE_2 int64 = 2 //pv
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
