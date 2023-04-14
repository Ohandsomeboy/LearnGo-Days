package main

import (
	"github.com/fogleman/gg"
)

func main() {
	// 创建一个宽500、高500的新画布
	dc := gg.NewContext(500, 500)

	// 绘制海绵宝宝的身体
	dc.SetRGB(253, 189, 9)             // 将颜色设置为黄色
	dc.DrawEllipse(250, 250, 150, 175) // 绘制一个椭圆形
	dc.Fill()

	// 绘制海绵宝宝的眼睛
	dc.SetRGB(255, 255, 255)    // 将颜色设置为白色
	dc.DrawCircle(200, 200, 30) // 绘制左眼
	dc.DrawCircle(300, 200, 30) // 绘制右眼
	dc.Fill()

	// 绘制海绵宝宝的瞳孔
	dc.SetRGB(0, 0, 0)          // 将颜色设置为黑色
	dc.DrawCircle(200, 200, 10) // 绘制左瞳孔
	dc.DrawCircle(300, 200, 10) // 绘制右瞳孔
	dc.Fill()

	// 绘制海绵宝宝的手脚
	dc.SetRGB(255, 255, 255)         // 将颜色设置为白色
	dc.DrawEllipse(125, 350, 50, 25) // 绘制左手
	dc.DrawEllipse(375, 350, 50, 25) // 绘制右手
	dc.DrawEllipse(200, 400, 50, 25) // 绘制左脚
	dc.DrawEllipse(300, 400, 50, 25) // 绘制右脚
	dc.Fill()

	// 保存绘制的图像
	dc.SavePNG("spongebob.png")
}
