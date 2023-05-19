package main

// 将摄氏度 ----> 开氏度 和 华氏度
func convertTemperature(celsius float64) []float64 {
	slice := []float64{celsius + 273.15, celsius*1.80 + 32.00}
	return slice
}
