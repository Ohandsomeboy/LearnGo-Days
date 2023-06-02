package 冒泡排序

func bubblesort(data []int) {
	var s = true
	for s {
		s = false
		for i := 1; i < len(data); i++ {
			if data[i-1] > data[i] {
				data[i], data[i-1] = data[i-1], data[i]
				s = true
			}
		}
	}
}
