//File  : server.go
//Author: duanhaobin
//Date  : 2020/5/27
/*
	自定义变量
	1.开发者可以在请求中自定义一些变量设置，自定义变量的获取优先级是最高的，可以覆盖原有的客户端提交参数。
	2.自定义变量往往也可以做请求流程的变量共享，但是需要注意的是该变量会成为请求参数的一部分，是对业务执行流程公开的变量。
	3.实例：使用jwt来授权时，通常会将部分用户信息(如userId)，放入到请求上下文中
*/
package main

import "fmt"

type Animal interface {
	MakeNoise()
}

type Dog struct {
	color string
}

/* Interface implementation */

func (d *Dog) MakeNoise() {
	fmt.Println("Bark!")
}

/* End Interface implementation */

func (d *Dog) WagTail() {
	fmt.Println(d.color + " dog: Wag wag")
}

func NewDog(color string) Animal {
	return &Dog{color}
}

func main() {
	dog := NewDog("Brown")
	dog.MakeNoise()
	//dog.WagTail()  报错
}
