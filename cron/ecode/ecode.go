package ecode

import "graduate/common/ecode"

var (
	InvalidOption = ecode.Register(1000, "无效的操作")
	InvalidSpec   = ecode.Register(1001, "无效的定时器")
)
