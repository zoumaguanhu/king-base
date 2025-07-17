package mq

// Topic
const (
	COMMAND_TOPIC     = "command"
	SYS_TOPIC         = "sys"
	COUPON_TOPIC      = "coupon"
	CART_TOPIC        = "cart"
	USER_ACTION_TOPIC = "user_action"
	ORDER_TOPIC       = "order"
)

// 队列名
const (
	COMMAND_CHANNEL     = "command_channel"
	SYS_CHANNEL         = "sys_channel"
	COUPON_CHANNEL      = "coupon_channel"
	CART_CHANNEL        = "cart_channel"
	USER_ACTION_CHANNEL = "user_action_channel"
	ORDER_CHANNEL       = "order_channel"
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
	ORDER_BZ       = "order_bz"
)

// CommandType类型
const (
	PRODUCT_UPDATE_COMMAND            = "product_update_command"
	BANNER_UPDATE_COMMAND             = "banner_update_command"
	SITE_UPDATE_COMMAND               = "site_update_command"
	RESET_PASSWD_COMMAND              = "reset_passwd_command"
	ISSUE_COUPON_COMMAND              = "issue_coupon_command"
	SIGN_UP_ISSUE_COUPON_COMMAND      = "sign_up_issue_coupon_command"
	ORDER_PAY_ISSUE_COUPON_COMMAND    = "order_pay_issue_coupon_command"
	ORDER_CREATE_COMMAND              = "order_create_command"
	ORDER_CANCEL_COMMAND              = "order_cancel_command"
	MAIL_CODE_COMMAND                 = "mail_code_command"
	CART_ADD_COMMAND                  = "cart_add_command"
	CART_EDIT_COMMAND                 = "cart_edit_command"
	CART_DEL_COMMAND                  = "cart_del_command"
	CART_REFRESH_COMMAND              = "cart_refresh_command"
	CART_BATCH_REMOVE_COMMAND         = "cart_batch_remove_command"
	CART_BATCH_FILL_ORDER_COMMAND     = "cart_batch_fill_order_command"
	CART_BATCH_RESTORE_COMMAND        = "cart_batch_restore_command"
	CART_CLEAR_COMMAND                = "cart_clear_command"
	USER_REGISTER_COMMAND             = "user_register_command"
	USER_SIGN_IN_COMMAND              = "user_sign_in_command"
	USER_SIGN_OUT_COMMAND             = "user_sign_out_command"
	USER_SIGN_UP_COMMAND              = "user_sign_up_command"
	USER_ADDRESS_REFRESH_COMMAND      = "user_address_refresh_command"
	USER_CART_REFRESH_COMMAND         = "user_cart_refresh_command"
	USER_ORDER_REFRESH_COMMAND        = "user_order_refresh_command"
	USER_ORDER_DETAIL_REFRESH_COMMAND = "user_order_detail_refresh_command"
	USER_COUPON_REFRESH_COMMAND       = "user_coupon_refresh_command"
	CONTENT_TP_UPDATE_COMMAND         = "content_tp_update_command"
)
