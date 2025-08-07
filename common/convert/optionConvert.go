package convert

import (
	"errors"
	"github.com/jinzhu/copier"
	"king.com/king/base/common/times"
	"time"
)

func Time2DefaultFormatStr() copier.Option {
	return Time2FormatStr(time.DateTime)
}
func Time2DefaultFormatAndId2Str() copier.Option {
	return Time2FormatStr(time.DateTime)
}
func Time2FormatStr(format string) copier.Option {
	option := copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
		Converters: []copier.TypeConverter{
			{
				SrcType: time.Time{},
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					s, ok := src.(time.Time)

					if !ok {
						return nil, errors.New("src type not matching")
					}

					return times.DateTimeFormat(s, format), nil
				},
			},
		},
	}

	return option
}
