//File  : server.go
//Author: duanhaobin
//Date  : 2020/5/26

/*
	http请求参数处理初识
	具体方法：https://godoc.org/github.com/gogf/gf/net/ghttp#Request
*/
package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

type RegisterReq struct {
	Name  string
	Pass  string `p:"password1"`
	Pass2 string `p:"password2"`
}

func main() {
	s := g.Server()
	s.BindHandler("/input", func(r *ghttp.Request) {
		// Get:常用方法，简化参数获取，GetRequest*的别名。不区分提交方式
		r.Response.Writeln(r.Get("amount"))
	})

	s.BindHandler("/query", func(r *ghttp.Request) {
		// GetQuery:获取GET方式传递过来的参数，包括Query String及Body参数解析
		r.Response.Writeln(r.GetQuery("amount"))
	})
	s.BindHandler("/form", func(r *ghttp.Request) {
		// GetForm:获取表单方式传递过来的参数，通常是POST请求；表单方式提交的参数Content-Type往往为
		// application/x-www-form-urlencoded, application/form-data, multipart/form-data, multipart/mixed等等。
		r.Response.Writeln(r.GetForm("amount"))
	})

	s.BindHandler("/struct", func(r *ghttp.Request) {
		// GetForm:获取表单方式传递过来的参数，通常是POST请求；表单方式提交的参数Content-Type往往为
		// application/x-www-form-urlencoded, application/form-data, multipart/form-data, multipart/mixed等等。
		var req *RegisterReq
		if err := r.GetStruct(&req); err != nil {
			r.Response.WritelnExit(err.Error())

		} else {
			glog.Info("测试 Exit1 .....")
			r.Response.WriteJsonExit(req)
		}
		glog.Info("测试 Exit2 .....")

	})

	s.BindHandler("/parse", func(r *ghttp.Request) {
		var req *RegisterReq
		// 从v1.11版本开始，推荐使用Parse方法来实现struct的转换，该方法是一个便捷方法，
		// 内部会自动进行转换及数据校验，但如果struct中没有校验tag的绑定将不会执行校验逻辑。
		if err := r.Parse(&req); err != nil {
			r.Response.WritelnExit(err.Error())

		} else {
			glog.Info("测试 Exit1 .....")
			r.Response.WriteJsonExit(req) // 执行完后程序就返回了，不会再执行后续的代码了。所以 '测试 Exit2'是不会输出的
		}
		glog.Info("测试 Exit2 .....")

	})
	s.Run()

}
