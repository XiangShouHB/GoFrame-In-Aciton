//File  : callback_server.go
//Author: duanhaobin
//Date  : 2020/5/22

/*
	函数注册
*/
package main

import (
	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//涉及到并发安全，因此total变量使用了gtype.Int并发安全类型
var total = gtype.NewInt()

// 实例一：包方法注册
func Total(r *ghttp.Request) {
	r.Response.Write("total:", total.Add(1))
}
func PkgMethodRegister(s *ghttp.Server) {
	s.BindHandler("/", Total)

}

// 实例二：对象方法注册
type Controller struct {
	// Int是用于int类型的并发安全操作的结构
	total *gtype.Int
}

func (c *Controller) Total(r *ghttp.Request) {
	r.Response.Write("total:", total.Add(1))
}
func ObjectMethodRegister(s *ghttp.Server) {
	// 加 '&' 是引用传递数据，结构体为值类型
	controller := &Controller{
		total: gtype.NewInt(),
	}
	s.BindHandler("/index", controller.Total)
}
func main() {
	/*
		回调函数注册
	*/
	s := g.Server()
	PkgMethodRegister(s)

	ObjectMethodRegister(s)
	s.Run()

}
