package constants

// 订单 字典常量
const (
	I18N   string = "i18n"
	LOCALE string = "locale"
	CN     string = "CN"
	OO     string = "OO"

	ORDER_STATUS string = "order_status"
)

var (
	VISIBLE_STATUS = []int64{1, 2, 3, 4, 6, 7, 8, 9}
	VALID_STATUS   = []int64{1, 2, 3, 4, 6, 7, 8, 9}
)

const (
	ORDER_SALE        = 1
	ORDER_PART_REFUND = 2
	ORDER_ALL_REFUND  = 3
)
const (
	EVENT_TYPE_1 = 1 //创建订单
	EVENT_TYPE_2 = 2 //用户支付完成
	EVENT_TYPE_3 = 3 //商户发货
	EVENT_TYPE_4 = 4 //已签收
	EVENT_TYPE_5 = 5 //已取消
	EVENT_TYPE_6 = 6 //部分退货
	EVENT_TYPE_7 = 7 //全部退货
	EVENT_TYPE_8 = 8 //退货完成
)
const (
	OPERATOR_TYPE_SYS      = 0 //系统
	OPERATOR_TYPE_MEMBER   = 1 //用户
	OPERATOR_TYPE_CUSTOMER = 2 //客服
	OPERATOR_TYPE_MERCHANT = 3 //商家
)
