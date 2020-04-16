package main

import "fmt"

func main() {
	// Q1 For-loop
	//1. 创建一个基于 for 的简单的循环。使其循环 10 次，并且使用 fmt 包打印出计数器的值

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 2. 用 goto 改写 1 的循环。关键字 for 不可使用。

	j := 0
LOOP:
	if j < 10 {
		println("LOOP==", j)
		j++
		goto LOOP
	}

	// 3.再次改写这个循环，使其遍历一个 array，并将这个 array 打印到屏幕上。
	ary := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for ar := range ary {
		fmt.Println("3.===", ar)
	}
}
