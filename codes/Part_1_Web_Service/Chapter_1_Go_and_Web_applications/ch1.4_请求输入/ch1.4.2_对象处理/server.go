//File  : server.go
//Author: duanhaobin
//Date  : 2020/5/27
/*
	对象处理.
	1.默认规则
		1).struct中需要匹配的属性必须为公开属性(首字母大写)d
		2).参数名称会自动按照 不区分大小写 且 忽略-/_/空格符号 的形式与struct属性进行匹配。
		3).如果匹配成功，那么将键值赋值给属性，如果无法匹配，那么忽略该键值
	2.自定义的参数映射规则可以通过为struct属性绑定tag实现，tag名称可以为p/param/params
	请求校验
*/

package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

// 使用p标签来指定该属性绑定的参数名称。password1参数将会映射到Pass1属性
// p 标签只是参数映射做了改变，不会修改返回json时的字段名
// v 标签是增加参数字段校验
// 注意：msg数要与rules数匹配，否则多余的规则会采用默认的信息来提示，即"字段值不合法"
type RegisterReq struct {
	Name  string `p:"username"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Pass  string `p:"password1" v:"required|length:6,30#请输入密码|密码长度不够"`
	Pass2 string `p:"password2" v:"required|length:6,30|same:password1#请确认密码|第二次输入的密码长度不够|两次密码不一致"`
}

// json 标签会修改返回json时的字段名
type RegisterRes struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func main() {
	s := g.Server()
	s.SetAccessLogEnabled(true)
	s.BindHandler("/register", func(r *ghttp.Request) {
		var req *RegisterReq
		// Parse 会先获取参数，然后自动校验数据
		// 有点鸡肋，不会返回具体错误。只是返回'字段值不合法'
		if err := r.Parse(&req); err != nil {
			r.Response.WriteJsonExit(RegisterRes{
				Code: 1,
				// 当请求校验错误时，所有校验失败的错误都返回了，这样对于用户体验不是特别友好
				Msg:  err.Error(),
				Data: nil,
			})
		}
		// 如果成功，返回请求参数，
		r.Response.WriteJson(RegisterRes{
			Code: 0,
			Msg:  "注册成功",
			Data: req,
		})
	})

	// 完善校验信息，返回关键点，而不是返回全部错误信息
	s.BindHandler("/register2", func(r *ghttp.Request) {
		var req *RegisterReq
		// Parse 会先获取参数，然后自动校验数据
		// 有点鸡肋，不会返回具体错误。只是返回'字段值不合法'
		if err := r.Parse(&req); err != nil {
			// Validation error.
			if v, ok := err.(*gvalid.Error); ok {
				r.Response.WriteJsonExit(RegisterRes{
					Code: 1,
					//返回FirstRule中得第一条规则错误信息
					Msg: v.FirstString(),
				})
			}
			// Other error.
			r.Response.WriteJsonExit(RegisterRes{
				Code: 1,
				// 当请求校验错误时，所有校验失败的错误都返回了，这样对于用户体验不是特别友好
				Msg: err.Error(),
			})
		}
		// 如果成功，返回请求参数，
		r.Response.WriteJson(RegisterRes{
			Code: 0,
			Msg:  "注册成功",
			Data: req,
		})
	})
	s.Run()
}
