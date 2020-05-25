//File  : 中间件设计.md
//Author: duanhaobin
//Date  : 2020/5/25

/*
	中间件设计实例
		1.CORS 跨域处理中间件
		2.鉴权处理中间件
*/
package main

import (
	"net/http"

	"github.com/gogf/gf/os/glog"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// CORS 中间件  前置中间件
func MiddleWareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Response.Writeln("cors in.....")
	r.Middleware.Next()
}

// 请求鉴权中间件  前置中间件
func MiddleWareAuth(r *ghttp.Request) {
	token := r.Get("token")
	if token == "mytoken" {
		r.Response.Writeln("auth in.....")
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

// 自定义日志处理  后置中间件
func MiddleWareLog(r *ghttp.Request) {
	r.Middleware.Next()
	errStr := ""
	if err := r.GetError(); err != nil {
		errStr = err.Error()
	}
	glog.Infof("请求状态:%d,URL:%s,err:%s", r.Response.Status, r.URL.Path, errStr)
}
func main() {
	s := g.Server()
	s.SetAccessLogEnabled(true)
	// 使用全局中间件
	s.Use(MiddleWareLog, MiddleWareCORS)

	s.Group("/api", func(group *ghttp.RouterGroup) {
		// 使用分组路由中间件
		group.Middleware(MiddleWareAuth)
		group.GET("/get", func(r *ghttp.Request) {
			r.Response.Writeln("get in ..............")
		})
	})
	s.Run()
}
