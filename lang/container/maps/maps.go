package main

import "fmt"

// map[K]V, map[K1]map[K2]V 格式
func main() {
	m := map[string]string{ //key在map里是无序的，是一个哈希map
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil,nil的值可以参与运算

	fmt.Println("m, m2, m3:")
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map m")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(`m["course"] =`, courseName)
	if causeName, ok := m["cause"]; ok { //如果输入错了的key依旧可以拿到值zero value
		fmt.Println(causeName)
	} else {
		fmt.Println("key 'cause' does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Printf("m[%q] before delete: %q, %v\n",
		"name", name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("m[%q] after delete: %q, %v\n",
		"name", name, ok)
}
