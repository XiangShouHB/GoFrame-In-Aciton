//File  : object_server01.go
//Author: duanhaobin
//Date  : 2020/5/22

package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type UserController struct{}

func (u *UserController) Index(r *ghttp.Request) {
	r.Response.Write("Index....")
}

func (u *UserController) Show(r *ghttp.Request) {
	r.Response.Write("Show....")
}

func main() {
	/*
		执行对象注册实例
	*/
	s := g.Server()
	// 绑定对象
	u := new(UserController)
	s.BindObject("/{.struct}/{.method}", u)
	s.Run()
}
