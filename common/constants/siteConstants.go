package constants

const (
	ITEM_TYPE_1 int64 = 1 //支付
	ITEM_TYPE_2 int64 = 2 //SMTP
	ITEM_TYPE_3 int64 = 3 //邮箱
	ITEM_TYPE_4 int64 = 4 //内容
)

// 支付方式
const (
	ITEM_PAY_PAY string = "PayPal"
)

// 邮件
const (
	SYS_MAIL      string = "1"
	CUSTOMER_MAIL string = "2"
)

// 内容
const (
	MAILE_TP int64 = 1
	BLOG     int64 = 2
)
