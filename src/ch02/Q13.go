package ch02

import "fmt"

// Q13. (1) 冒泡排序
//1. 编写一个针对 int 类型的 slice 冒泡排序的函数。
//这里它在一个列表上重复步骤来排序，比较每个相䩪的元素，并且顺序错误的时候，交换它们。
//一遍一遍扫描列表，直到没有交换为止，这意味着列表排序完成。算法得名于更小的元素就像 “泡泡” 一样冒到列表的別端

func Q13() {
	fmt.Println("=========Q13=========")
	sls := []int{1, 21, 13, 5, 71, 8, 9, 10, 11, 13}
	fmt.Println("排序前 ", sls)
	t131(sls)
	fmt.Println("排序后 ", sls)
}

func t131(sls []int) {

	for i := 0; i < len(sls)-1; i++ {
		for j := i + 1; j < len(sls); j++ {
			if sls[i] > sls[j] {
				sls[i], sls[j] = sls[j], sls[i]
			}
		}
	}
}
