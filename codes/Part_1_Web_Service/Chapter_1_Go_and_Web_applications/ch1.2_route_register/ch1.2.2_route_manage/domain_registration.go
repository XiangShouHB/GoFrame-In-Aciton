//File  : domain_registration.go
//Author: duanhaobin
//Date  : 2020/5/22

package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	/*
		域名注册实例
	*/
	d := g.Server().Domain("localhost")
	d.BindHandler("/index", func(r *ghttp.Request) {
		r.Response.Writef("域名注册:%v", r.Router)
	})
	g.Server().Run()
}
