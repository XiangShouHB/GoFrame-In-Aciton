//File  : sample_demo.go
//Author: duanhaobin
//Date  : 2020/5/21

package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	// 动态路由原则之命名匹配规则 :name
	s.BindHandler("/:name", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})
	s.BindHandler("/:name/update", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})
	s.BindHandler("/:name/:action", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})
	// 动态路由原则之模糊匹配规则 *any
	s.BindHandler("/:name/*any", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})
	// 动态路由原则之字段匹配原则 {field}
	s.BindHandler("/user/list/{field}.html", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})
	s.SetPort(8199)
	s.Run()
}
