//File  : router_rules.go
//Author: duanhaobin
//Date  : 2020/5/19

package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {

	// 创建一个默认的 Server 对象
	s := g.Server()

	// 该路由规则仅会在GET请求下有效
	s.BindHandler("GET:/index", func(r *ghttp.Request) {
		r.Response.WriteJson(r.Router) // r.Router是当前匹配的路由规则信息
	})

	// 该路由规则仅会在GET请求及localhost域名下有效
	// http://127.0.0.1:8090/home 是无效的,必须是： http://home:8090/home
	s.BindHandler("GET:/home@localhost", func(r *ghttp.Request) {
		r.Response.WriteJson(r.Router)
	})

	// 动态路由实例一：分页路由示例
	s.BindHandler("/user/list/{page}.html", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router, r.Get("page"))
	})

	// 动态路由实例二：{xxx} 规则与 :xxx 规则混合使用
	// http://127.0.0.1:8090/prod/list/test.html
	s.BindHandler("/{object}/:attr/{act}.html", func(r *ghttp.Request) {
		r.Response.Writeln("object->", r.Get("object")) // object->prod
		r.Response.Writeln("attr->", r.Get("attr"))     // attr->attr
		r.Response.Writeln("act->", r.Get("act"))       // act->act
	})
	// 指定端口:8090;默认为:80
	s.SetPort(8090)
	s.Run()
}
