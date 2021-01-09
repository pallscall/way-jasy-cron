package ecode

import "graduate/common/ecode"

var (
	InvalidUser       = ecode.Register(2000, "用户不存在")
	InvalidPassword   = ecode.Register(2001, "密码错误")
	UserExist         = ecode.Register(2002, "用户名已存在")
)
