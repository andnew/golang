package ch02

import "fmt"

//Q14. (1) 函数返回一个函数
//1. 编写一个函数返回另一个函数，返回的函数的作用是对一个整数 +2。函数的名称叫做 plusTwo。然后可以像下面这样使用：
//p := plusTwo()
//fmt.Printf("%v\n", p(2))
//应该打印 4。参阅第 31 页的 “回调” 小节了解更多相关信息。
//2. 使 1 中的函数更加通用化，创建一个 plusX(x) 函数，返回一个函数用于对整数加上 x。

func Q14() {
	fmt.Println("=========Q14=========")
	t141()
}

func t141() {
	p := plusTwo()
	fmt.Println(p(2))
}

func plusTwo() func(v int) int {
	return func(v int) int {
		return v + 2
	}
}
