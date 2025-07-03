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
	VISIBLE_STATUS = [...]int64{1, 2, 3, 4, 6, 7, 8, 9}
)

const (
	ORDER_SALE        = 1
	ORDER_PART_REFUND = 2
	ORDER_ALL_REFUND  = 3
)
