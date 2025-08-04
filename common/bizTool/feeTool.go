package bizTool

import (
	"king.com/king/base/common/constants"
	"king.com/king/base/common/db/rdb"
	"king.com/king/base/common/strs"
)

func FeePrice(price string, fe *rdb.FeeInfo) string {
	switch fe.FeeType {
	case constants.ITEM_KEY_1:
		return constants.ZERO_STR
	case constants.ITEM_KEY_2:
		return fe.FeeVal
	case constants.ITEM_KEY_3: //此刻的price是订单中该类型的订单总价
		if strs.StrToFloat64(price)-strs.ConvertStr2Float(fe.FeeOption) > 0 {
			return constants.ZERO_STR
		}
		return fe.FeeVal
	case constants.ITEM_KEY_4:
		fp := strs.StrToFloat64(price) * strs.ConvertStr2Float(fe.FeeVal) / 100
		return strs.ConvertFloat2Str(fp)

	}
	return constants.ZERO_STR
}
