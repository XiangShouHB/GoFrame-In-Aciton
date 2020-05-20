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
		r.Response.Writeln("This is a index page!")
		r.Response.WriteJson(r.Router)
	})

	// 该路由规则仅会在GET请求及localhost域名下有效
	// http://127.0.0.1:8090/home 是无效的
	s.BindHandler("GET:/home@localhost", func(r *ghttp.Request) {
		r.Response.Writeln("This is a home page which valids under localhost domain name ")
		r.Response.WriteJson(r.Router)
	})
	// 指定端口:8090;默认为:80
	s.SetPort(8090)
	s.Run()
}
