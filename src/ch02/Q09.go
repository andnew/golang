package ch02

import "fmt"

//Q9. (1) 变参

func Q9() {
	fmt.Println("=========Q09=========")
	t91([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

//1. 编写函数接受整数类型变参，并且每行打印一个数字。
func t91(vars []int) {
	for _, v := range vars {
		fmt.Println(v)
	}
}
