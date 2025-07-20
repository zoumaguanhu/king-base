package constants

const UN_KNOW string = "N/A"
const DEFAULT_STR string = ""
const HTTPS_PREFIX string = "https"
const EMPTY_STR string = " "
const SEPARATOR_STR string = ";"
const SEP_STR string = ","
const BEARER string = "Bearer"
const AUTHORIZATION string = "Authorization"
const ZERO_STR string = "0"
const FIRST_STR string = "1"
const ZERO_INT64 int64 = 0
const ONE_INT64 int64 = 1
const TEN_INT64 int64 = 10
const TEN_INT int = 10
const NUM_120_INT int = 120
const NUM_120_FLOAT64 float64 = 120
const TRUE bool = true
const FALSE bool = false
const LETTER_BYTES = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const TRACE_ID string = "traceId"
const USER_ID string = "userID"
const FINGER_FLG string = "finger"
const FINGER_FILL string = "|jwt@hex@client|"
const ROLE string = "role"
const VIRT_ID string = "virtId"
const VIRT_IDS string = "virtIds"
const SITE_ROLE string = "sr"
const HOST_FLG string = "hostFlg"
const TEN_THOUSAND_64 int64 = 10000
const TEN_THOUSAND int = 10000
const DELIVERY_FEE string = "delivery_fee"
const VISIBLE_1 int64 = 1
const VISIBLE_2 int64 = 2
const PER_CENT string = "%"
const DLX_PREFIX string = "dlx_"
const DIRECT string = "direct"
const SHORT_LINE = "_"
const X_FORWARDED_FOR = "X-Forwarded-For"
const X_REAL_IP = "X-Real-IP"
const (
	FIELD_VIRT_ID  = "virt_id"
	NOT_FOUNT_SITE = "not found site"
)
const ZERO_ZERO = "00"
const DOT = "."
const RD_SEPARATOR = ":"
const DATE_366 int = 366
const CG_FLG string = "cgFlg"
const X_CODE = "xCode"

// 双引号空
const EMPTY_STRING = ""
const PAYLOAD = "payload"
const MD5 = "md5"
const EXP = "exp"
const IAT = "iat"
const (
	PRODUCT_SKU_ERROR = 300111
)
const (
	P_LEVEL_1 = 1
	P_LEVEL_2 = 2
	P_LEVEL_3 = 3
)
const (
	STATUS_1 int64 = 1
	STATUS_2 int64 = 2
	STATUS_3 int64 = 3
	STATUS_4 int64 = 4
	STATUS_5 int64 = 5
	STATUS_6 int64 = 6
	STATUS_7 int64 = 7
	STATUS_8 int64 = 8
)
const (
	PUBLISH_1 int64 = 1
	PUBLISH_2 int64 = 2
)
const (
	HTTP_CONTENT_TYPE      string = "Content-Type"
	HTTP_APPLICATION_JSON  string = "application/json"
	HTTP_APPLICATION_X_WWW string = "application/x-www-form-urlencoded"
	HTTP_AUTHORIZATION     string = "Authorization"

	HTTP_TIMEOUT   int64   = 30000
	HTTP_TIMEOUT_F float64 = 30000
	HTTP_OK        int     = 200
	HTTP_201       int     = 201
	HTTP_401       int     = 401
	HTTP_NO_BODY   string  = ""
)

const (
	PAYPAL                     string = "PayPal"
	PAYPAI_GRANT_TYPE          string = "grant_type"
	PAYPAI_CLIENT_CREDENTIALS  string = "client_credentials"
	PAYPAL_REQUEST_ID          string = "PayPal-Request-Id"
	REF_PAYER_ACTION           string = "payer-action"
	USD                        string = "USD"
	DOLLAR                     string = "$"
	EN_US                      string = "en-US"
	IMMEDIATE_PAYMENT_REQUIRED string = "IMMEDIATE_PAYMENT_REQUIRED"
	LANDING_PAGE               string = "LOGIN"
	SET_PROVIDED_ADDRESS       string = "SET_PROVIDED_ADDRESS"
	PAY_NOW                    string = "PAY_NOW"
	ORDER_EXPIRE               int64  = 3
	DIGITAL_GOODS              string = "DIGITAL_GOODS"
	PHYSICAL_GOODS             string = "PHYSICAL_GOODS"
	DONATION                   string = "DONATION"
)

const (
	CAPTURE   = "CAPTURE"
	COMPLETED = "COMPLETED"
)
const (
	P_TYPE_1 int64 = 1
	P_TYPE_2 int64 = 2
	P_TYPE_3 int64 = 3
	P_TYPE_4 int64 = 4
)
const (
	ROLE_0 int64 = 0
	ROLE_1 int64 = 1
	ROLE_2 int64 = 2
	ROLE_3 int64 = 3
)
const (
	COMMAND_TYPE_UPDATE = "update"
)

type Virt struct {
	VirtId int64 `json:"virtId"`
}

type HeaderOption struct {
	Channel    string `json:"channel"`
	Platform   string `json:"platform"`
	PCode      string `json:"PCode"`
	AcCode     string `json:"acCode"`
	InviteCode string `json:"inviteCode"`
}

const (
	REFRESH_CONTURY      = "refresh_country"
	REFRESH_PAY_METHOD   = "refresh_pay_method"
	REFRESH_LOCATION_API = "refresh_location_api"
)
