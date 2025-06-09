package mq

// 队列名
const (
	COMMAND_QUEUE = "command_queue"
	SYS_QUEUE     = "sys_queue"
	COUPON_QUEUE  = "coupon_queue"
	CART_QUEUE    = "cart_queue"
)

// bz类型
const (
	PRODUCT_BZ = "product_bz"
	SITE_BZ    = "site"
	BANNER_BZ  = "banner_bz"
	SYS_BZ     = "sys_bz"
	COUPONE_BZ = "coupon_bz"
	CART_BZ    = "cart_bz"
)

// CommandType类型
const (
	COMMAND_TYPE_UPDATE  = "update_command"
	RESET_PASSWD_COMMAND = "reset_passwd_command"
	ISSUE_COUPON_COMMAND = "issue_coupon_command"
	MAIL_CODE_COMMAND    = "mail_code_command"
	CART_ADD_COMMAND     = "cart_add_command"
	CART_EDIT_COMMAND    = "cart_edit_command"
	CART_DEL_COMMAND     = "cart_del_command"
)
