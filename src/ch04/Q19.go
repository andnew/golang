package ch04

import "fmt"

//Q19. (1) 指针
//1. 假设定义了下面的结构：
//type Person struct {
//name string
//age int
//}
//下面两行之间的区别是什么？
//var p1 Person
//p2 := new(Person)
//2. 下面两个内存分配的区别是什么？
//func Set(t *T) {
//x = t
//} 和
//func Set(t T) {
//x= &t
//}

func Q19() {
	var p1 Person     //var p1 Person 分配了 Person-值 给 p1。 p1 的类型是 Person。
	p2 := new(Person) //p2 := new(Person) 分配了内存并且将指针赋值给 p2。 p2 的类型是*Person。
	fmt.Printf("p1 [%v]\n", p1)
	fmt.Printf("p2 [%v]\n", p2)
}

type Person struct {
	name string
	age  int
}

//func Set(t *T) {
//	x = t //x 指向了 t 指向的内容，也就是实际上的参数指向的内容。
//}
//func Set(t T) {
//	x = &t // x 指向一个新的（堆上分配的）变量 t，其包含了实际参数值的副本。
//}

//在第二个函数中， x 指向一个新的（堆上分配的）变量 t，其包含了实际参数值的副本。
//在第一个函数中， x 指向了 t 指向的内容，也就是实际上的参数指向的内容。
//因此在第二个函数，我们有了 “额外” 的变量存储了相关值的副本。
