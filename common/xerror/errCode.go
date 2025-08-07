package xerror

// 成功返回
const OK int = 200

// 业务码

const (
	SERVER_ERROR             = 20000 //系统异常
	SERVER_COPIER_COPY       = 20001 //copier组件异常
	TOKEN_GENERATE_ERROR     = 20002 //token生成失败
	TOKEN_PARSE_ERROR        = 20003 //token解析失败
	TOKEN_NOT_VALID_ERROR    = 20004 //token无效
	SERVER_STR_TO_TIME_ERROR = 20005 //字符串转Time无效

)
const (
	SERVER_QUERY_ERROR        = 200100 //查询异常
	SERVER_QUERY_TOTAL_ERROR  = 200101 //查询total异常
	SERVER_QUERY_DETAIL_ERROR = 200102 //查询detail异常
	SERVER_UPDATE_ERROR       = 200103 //执行更新失败
	THIRD_ERROR               = 200104 //执行更新失败
)
