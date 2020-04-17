package ch02

import "fmt"

//Q8. (1) 栈
//1. 创建一个固定大小保存整数的栈。它无须超出限制的增长。定义 push 函数——将数据放入栈，和 pop 函数——从栈中取得内容。栈应当是后进先出（ LIFO）的。
//Figure 2.1. 一个简单的 LIFO 栈
//push(k)
//i k pop() k
//l m
//i++
//i--
//2. 更进一步。编写一个 String 方法将栈转化为字符串形式的表达。可以这样的
//方式打印整个栈： fmt.Printf("My stack %v\n", stack)
//栈可以被输出成这样的形式： [0:m] [1:l] [2:k]

func Q8() {
	fmt.Println("=========Q08=========")
	t81()
}

type stack struct {
	k   int     //栈的长度
	ary [10]int //栈的内存
}

func (s *stack) push(v int) {
	if s.k > 9 {
		return
	} else {
		s.ary[s.k] = v
		s.k = s.k + 1
	}
}
func (s *stack) pop() int {
	if s.k == 0 {
		return -1
	} else {
		res := s.ary[s.k-1]
		s.k = s.k - 1
		return res
	}
}

func (s *stack) String() string {
	scont := ""
	for i := 0; i < s.k-1; i++ {
		scont += fmt.Sprintf("[%d:%d]", i, s.ary[i])
		if i < s.k-2 {
			scont += ","
		}
	}
	return scont
}

func t81() {
	var s stack
	s.push(1)
	s.push(2)
	fmt.Println(s.k, s.ary)
	fmt.Println(s.pop())
	fmt.Println(s.pop())

	var s2 stack
	s2.push(1)
	s2.push(2)
	s2.push(3)
	s2.push(5)
	fmt.Println(s2.String())
}
