package ch02

import "fmt"

//Q12. (0) 最小值和最大值
//1. 编写一个函数，找到 int slice ([]int) 中的最大值。
//2. 编写一个函数，找到 int slice ([]int) 中的最小值。

func Q12() {
	fmt.Println("=========Q12=========")
	sls := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	max := t121(sls)
	fmt.Println("slice ", sls, " , 最大值 ", max)
	min := t122(sls)
	fmt.Println("slice ", sls, " , 最大值 ", min)

}

func t121(slis []int) int {

	if len(slis) == 0 {
		fmt.Println("没有最大值")
		return -1
	}

	max := slis[0]

	for i := 1; i < len(slis); i++ {
		if max < slis[i] {
			max = slis[i]
		}
	}
	return max
}

func t122(slis []int) int {

	if len(slis) == 0 {
		fmt.Println("没有最大值")
		return -1
	}

	min := slis[0]

	for i := 1; i < len(slis); i++ {
		if min > slis[i] {
			min = slis[i]
		}
	}
	return min
}
