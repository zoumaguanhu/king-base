package constants

const DEFAULT_STR string = ""
const SEPARATOR_STR string = ";"
const ZERO_STR string = "0"
const ZERO_INT64 int64 = 0
const TEN_INT64 int64 = 10
const LETTER_BYTES = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const TRACE_ID string = "TRACE_ID"
const USER_ID string = "userID"
const ROLE string = "role"
const VIRT_ID string = "virtId"
const VIRT_FLG string = "virtFlg"
const TEN_THOUSAND_64 int64 = 10000
const TEN_THOUSAND int = 10000
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
	HTTP_CONTENT_TYPE      string  = "Content-Type"
	HTTP_APPLICATION_JSON  string  = "application/json"
	HTTP_APPLICATION_X_WWW string  = "application/x-www-form-urlencoded"
	HTTP_AUTHORIZATION     string  = "Authorization"
	BEARER                 string  = "Bearer"
	HTTP_TIMEOUT           int64   = 30000
	HTTP_TIMEOUT_F         float64 = 30000
	HTTP_OK                int     = 200
	HTTP_201               int     = 201
	HTTP_401               int     = 401
)
const (
	PAYPAL                     string = "PayPal"
	PAYPAI_GRANT_TYPE          string = "grant_type"
	PAYPAI_CLIENT_CREDENTIALS  string = "client_credentials"
	PAYPAL_REQUEST_ID          string = "PayPal-Request-Id"
	REF_PAYER_ACTION           string = "payer-action"
	USD                        string = "USD"
	EN_US                      string = "en-US"
	IMMEDIATE_PAYMENT_REQUIRED string = "IMMEDIATE_PAYMENT_REQUIRED"
	LANDING_PAGE               string = "LOGIN"
	SET_PROVIDED_ADDRESS       string = "SET_PROVIDED_ADDRESS"
	PAY_NOW                    string = "PAY_NOW"
)

const (
	CAPTURE = "CAPTURE"
)
