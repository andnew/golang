package ch02

import "fmt"

//Q6. (0) 整数顺序
//1. 编写函数，返回其（两个）参数正确的（自然）数字顺序：
//f(7,2) ! 2,7
//f(2,7) ! 2,7

func Q6() {
	fmt.Println("=========Q06=========")
	num1, num2 := t61(7, 2)
	fmt.Println(num1, num2)
	num1, num2 = t61(2, 7)
	fmt.Println(num1, num2)
}

func t61(num1, num2 int) (int, int) {

	if num1 > num2 {
		return num2, num1
	} else {
		return num1, num2
	}
}
