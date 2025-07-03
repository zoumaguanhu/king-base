package mq

// Topic
const (
	COMMAND_TOPIC     = "command"
	SYS_TOPIC         = "sys"
	COUPON_TOPIC      = "coupon"
	CART_TOPIC        = "cart"
	USER_ACTION_TOPIC = "user_action"
)

// 队列名
const (
	COMMAND_CHANNEL     = "command_channel"
	SYS_CHANNEL         = "sys_channel"
	COUPON_CHANNEL      = "coupon_channel"
	CART_CHANNEL        = "cart_channel"
	USER_ACTION_CHANNEL = "user_action_channel"
)

// bz类型

const (
	PRODUCT_BZ     = "product_bz"
	SITE_BZ        = "site_bz"
	BANNER_BZ      = "banner_bz"
	SYS_BZ         = "sys_bz"
	COUPONE_BZ     = "coupon_bz"
	CART_BZ        = "cart_bz"
	USER_ACTION_BZ = "user_action_bz"
	CONTENT_TP_BZ  = "content_tp_bz"
)

// CommandType类型
const (
	PRODUCT_UPDATE_COMMAND         = "product_update_command"
	BANNER_UPDATE_COMMAND          = "banner_update_command"
	SITE_UPDATE_COMMAND            = "site_update_command"
	RESET_PASSWD_COMMAND           = "reset_passwd_command"
	ISSUE_COUPON_COMMAND           = "issue_coupon_command"
	SIGN_UP_ISSUE_COUPON_COMMAND   = "sign_up_issue_coupon_command"
	ORDER_PAY_ISSUE_COUPON_COMMAND = "order_pay_issue_coupon_command"
	MAIL_CODE_COMMAND              = "mail_code_command"
	CART_ADD_COMMAND               = "cart_add_command"
	CART_EDIT_COMMAND              = "cart_edit_command"
	CART_DEL_COMMAND               = "cart_del_command"
	CART_REFRESH_COMMAND           = "cart_refresh_command"
	USER_REGISTER_COMMAND          = "user_register_command"
	USER_SIGN_IN_COMMAND           = "user_sign_in_command"
	USER_SIGN_UP_COMMAND           = "user_sign_up_command"
	CONTENT_TP_UPDATE_COMMAND      = "content_tp_update_command"
)
