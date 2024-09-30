package xerror

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"king.com/king/rpc-ums/pkg/constants"
	"strings"
)

var message map[int]string

func init() {
	message = make(map[int]string)
	message[OK] = "SUCCESS"
	message[SERVER_COPIER_COPY] = "COPIER组件异常"
	message[SERVER_QUERY_ERROR] = "数据库查询异常"
	message[SERVER_QUERY_TOTAL_ERROR] = "数据库查询总数异常"
	message[SERVER_QUERY_DETAIL_ERROR] = "数据库查询详情异常"
	message[TOKEN_GENERATE_ERROR] = "token生成错误"
	message[TOKEN_PARSE_ERROR] = "token解析错误"
	message[TOKEN_NOT_VALID_ERROR] = "token无效"
	message[SERVER_STR_TO_TIME_ERROR] = "字符串转time无效"
	message[SERVER_UPDATE_ERROR] = "执行更新失败"

}

// 根据错误码设置返回错误信息
func New(errCode codes.Code, pd ...any) (err error) {
	return status.Error(errCode, FinalMsg(int(errCode), pd...))
}

// 返回消息内容（支持内容格式化）
func FinalMsg(errCode int, pd ...any) string {
	if msg, ok := message[errCode]; ok {
		if pd != nil {
			return strings.TrimRight(fmt.Sprintf(msg, pd...), constants.COLON)
		}
		return strings.TrimRight(msg, constants.ERROR_SUFFIX)
	}
	//返回默认错误信息
	return message[SERVER_ERROR]
}

func IsCodeErr(errCode int) bool {
	if _, ok := message[errCode]; ok {
		return true
	} else {
		return false
	}
}
