package ch04

import "fmt"

// Q17. (1) 指针运算
//1. 在正文的第 54 页有这样的文字：
//…这里没有指针运算，因此如果这样写： *p++，它被解释为 (*p)++：
//首先解析引用然后增加值。
//当像这样增加一个值的时候，什么类型可以工作？
//2. 为什么它不能工作在所有类型上？

func Q17() {
	fmt.Println("=========Q17=========")
	p := 100
	q := &p
	var p1 uint = 12

	fmt.Printf("p  int %v\n", *q)
	fmt.Printf("p1 int %v\n", &p1)

	var i int
	var n *int
	n = &i
	*n++
	fmt.Printf("%v\n", *n)
}

//这仅能工作于指向数字（ int, uint 等等）的指针值。
