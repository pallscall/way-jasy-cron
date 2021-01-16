// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api2.proto

/*
Package userapi is a generated blademaster stub package.
This code was generated with kratos/tool/protobuf/protoc-gen-bm v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:
	api2.proto
*/
package userapi

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathUserLoginUrl = "/user/login"

// UserBMServer is the server API for User service.
type UserBMServer interface {
	LoginUrl(ctx context.Context, req *LoginReq) (resp *LoginResp, err error)
}

var UserSvc UserBMServer

func userLoginUrl(c *bm.Context) {
	p := new(LoginReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := UserSvc.LoginUrl(c, p)
	c.JSON(resp, err)
}

// RegisterUserBMServer Register the blademaster route
func RegisterUserBMServer(e *bm.Engine, server UserBMServer) {
	UserSvc = server
	e.GET("/user/login", userLoginUrl)
}
