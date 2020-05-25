//File  : group_router.go
//Author: duanhaobin
//Date  : 2020/5/25

/*
	分组路由及层架路由
*/
package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	// 成功也启动日志输出
	g.Server().SetAccessLogEnabled(true)
	// 定义分组路由
	s.Group("/api", func(group *ghttp.RouterGroup) {
		// 设置层级路由
		group.Group("/user", func(group *ghttp.RouterGroup) {
			group.GET("/list", func(r *ghttp.Request) {
				r.Response.Write("get user list....")
			})
			group.POST("/add", func(r *ghttp.Request) {
				r.Response.Write("add new user....")
			})
		})
		// 设置层级路由
		group.Group("/order", func(group *ghttp.RouterGroup) {
			group.GET("/list", func(r *ghttp.Request) {
				r.Response.Write("get order list....")
			})
			group.POST("/add", func(r *ghttp.Request) {
				r.Response.Write("add new order....")
			})
		})
	})
	s.Run()
}
