package constants

var statMap map[int64]string

const (
	SALE_TOTAL             = 11
	SALE_AMOUNT            = 12
	SALE_7_TOTAL           = 13
	SALE_7_AMOUNT          = 14
	STATS_UN_PAY           = 21
	STATS_UN_SEND          = 22
	STATS_UN_RECEIVE       = 23
	STATS_REFUNDING        = 24
	OTHER_UN_PUBLISH       = 31
	OTHER_STOCK_UP         = 32
	OTHER_ON_SALE          = 33
	OTHER_TAKE_OFF_PRODUCT = 34
)

func init() {
	statMap = make(map[int64]string)
	statMap[SALE_TOTAL] = "今日销售量"
	statMap[SALE_7_TOTAL] = "近7日销售量"
	statMap[SALE_AMOUNT] = "今日销售金额($)"
	statMap[SALE_7_AMOUNT] = "近7日销售金额($)"
	statMap[STATS_UN_PAY] = "未支付"
	statMap[STATS_UN_SEND] = "待发货"
	statMap[STATS_UN_RECEIVE] = "待收货"
	statMap[STATS_REFUNDING] = "退货申请"
	statMap[OTHER_UN_PUBLISH] = "未发布"
	statMap[OTHER_STOCK_UP] = "备货建议"
	statMap[OTHER_ON_SALE] = "在售商品"
	statMap[OTHER_TAKE_OFF_PRODUCT] = "已下架商品"
}
func FindName(k int64) string {
	if v, ok := statMap[k]; ok {
		return v
	}
	return DEFAULT_STR
}
