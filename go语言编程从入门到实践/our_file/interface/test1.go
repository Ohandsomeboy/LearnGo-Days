package main

import (
	"fmt"
)

// 定义接口
type actions interface {
	walk()
	runs() (int, int)
	speak(content string, speed int)
	rest(sleepTime int) int
}

type cats struct {
	name string
}

func (c *cats) walk() {
	fmt.Printf("%vd在散步\n", c.name)
}

func (c *cats) runs() (int, int) {
	fmt.Printf("%v在跑步\n", c.name)
	speed := 10
	time := 1
	return speed, time
}

func (c *cats) speak(content string, speed int) {
	fmt.Printf("%v在说话：%v,语速：%v\n", c.name, content, speed)
}

func (c *cats) rest(sleepTime int) int {
	fmt.Printf("%v在休息, 入睡时间：%v小时\n", c.name, sleepTime)
	return sleepTime
}

func main() {
	//定义接口变量
	var a actions
	//结构体实例化
	c := cats{name: "Kitty"}
	//结构体实例化变量的指针赋值给接口变量
	a = &c
	//调用接口中的方法
	a.walk()
	speed, time := a.runs()
	fmt.Printf("跑步速度: %v, 跑步时间：%v\n", speed, time)
	a.speak("喵喵", 2)
	sleepTime := a.rest(10)
	fmt.Printf("入睡时间：%v小时\n", sleepTime)
}

//综上所述，接口的定义与使用总结如下:
//1)接口是使用关键字type和interface定义的，接口方法只需设置方法名称、参数及其数据类型、返回值的数据类型。
//2)接口方法的功能逻辑由结构体方法实现，接口无法单独使用，它必须与结构体组合使用。
//3)使用接口必须创建接口变量和实例化结构体，然后将结构体实例化变量或变量的内存地址赋值给接口变量，完成结构体与接口的绑定。
//4)接口变量只能调用接口中定义的方法，结构体实例化变量不仅能调用接口方法，还能调用接口之外的结构体方法和结构体成员。
//5)如果结构体绑定了接口，结构体必须为接口中的每个方法定义相应的结构体方法，否则程序提示as some methods are missing异常。
