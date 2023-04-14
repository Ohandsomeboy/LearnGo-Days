package main

import "fmt"

//鸭子类型：将接口与结构体的绑定过程以函数实现，
//只要传入结构体实例化变量就能自动执行接口方法

type actions interface {
	speak(content string)
}

type duck struct {
	name string
}

type cat struct {
	name string
}

func (d *duck) speak(content string) {
	fmt.Printf("%v\n", d.name)
}

func (c *cat) speak(content string) {
	fmt.Printf("%v在说话:%v\n", c.name, content)
}

func speaking(a actions, content string) {
	a.speak(content)
}

func main() {
	//实例化结构体
	d := duck{name: "唐老鸭"}
	c := cat{name: "凯蒂猫"}
	//调用函数
	speaking(&d, "嘎嘎")
	speaking(&c, "喵喵")
}
