package ch02

import "fmt"

//Q10. (1) 斐波那契
//1. 斐波那契数列以： 1; 1; 2; 3; 5; 8; 13; : : : 开始。或者用数学形式表达： x1 = 1; x2 = 1; xn = xn−1 + xn−2 8n > 2。
//编写一个接受 int 值的函数，并给出这个值得到的斐波那契数列
func Q10() {
	fmt.Println("=========Q10=========")
	fmt.Println("输入是1 ", t101(1))
	fmt.Println("输入是2 ", t101(2))
	fmt.Println("输入是3 ", t101(3))
	fmt.Println("输入是7 ", t101(7))
}

func t101(xn int) []int {
	ary := []int{1, 1}
	if xn >= 1 && xn <= 2 {
		return ary[0:xn]
	}

	for i := 3; i <= xn; i++ {
		v := ary[i-2] + ary[i-3]
		ary = append(ary, v)
	}
	return ary
}
