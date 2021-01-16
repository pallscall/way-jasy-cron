// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api2.proto

/*
Package api is a generated blademaster stub package.
This code was generated with kratos/tool/protobuf/protoc-gen-bm v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:
	api2.proto
*/
package api

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathJobLogin = "/demo.service.v1.Job/Login"

// JobBMServer is the server API for Job service.
type JobBMServer interface {
	Login(ctx context.Context, req *LoginReq) (resp *LoginResp, err error)
}

var JobSvc JobBMServer

func jobLogin(c *bm.Context) {
	p := new(LoginReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := JobSvc.Login(c, p)
	c.JSON(resp, err)
}

// RegisterJobBMServer Register the blademaster route
func RegisterJobBMServer(e *bm.Engine, server JobBMServer) {
	JobSvc = server
	e.GET("/demo.service.v1.Job/Login", jobLogin)
}
