package main

import "fmt"

func main() {
	// 定义一个数组
	arrayList := []string{
		"I",
		"am",
		"stupid",
		"and",
		"weak",
	}

	// 查看原始的数据
	fmt.Println("Old: ", arrayList)
	// 查看每个下标对应的数据
	for i, v := range arrayList {
		fmt.Println(i, v)
	}

	//	使用 for 循环更新部分字段
	for i, v := range arrayList {
		if v == "stupid" {
			arrayList[i] = "smart"
		} else if v == "weak" {
			arrayList[i] = "strong"
		}
	}

	// 查看更新后的的数据
	fmt.Println("New: ", arrayList)
}
