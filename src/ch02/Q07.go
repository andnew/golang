package ch02

import "fmt"

//Q7. (1) 作用域
//1. 下面的程序有什么错误？
func Q7() {
	fmt.Println("=========Q07=========")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}
	//fmt.Printf("%v\n", i) // i 的作用域只能在for 循环内
}
