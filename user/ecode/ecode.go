package ecode

import "way-jasy-cron/common/ecode"

var (
	InLegalOpt        = ecode.Register(1500, "非法操作")

	InvalidUser       = ecode.Register(2000, "用户不存在")
	InvalidPassword   = ecode.Register(2001, "密码错误")
	UserExist         = ecode.Register(2002, "用户名已存在")

)
