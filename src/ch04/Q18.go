package ch04

import "fmt"

//Q18. (2) 使用 interface 的 map 函数
//1. 使用练习 Q11 的答案，利用 interface 使其更加通用。让它至少能同时工作于 int 和 string。

func Q18() {
	fmt.Println("=========Q18=========")
	m := []e{1, 2, 3, 4}
	s := []e{"a", "b", "c", "d"}
	mf := Map(m, func2)
	sf := Map(s, func2)
	fmt.Printf("%v\n", mf)
	fmt.Printf("%v\n", sf)
}

type e interface{}

func Map(n []e, f func(e) e) []e {
	m := make([]e, len(n))
	for k, v := range n {
		m[k] = f(v)
	}
	return m
}

//
//func Map(f func(e) e, l []e) []e {
//	j := make([]e, len(l))
//	for k, v := range l {
//		j[k] = f(v)
//	}
//	return j
//}

func func2(f e) e {
	switch f.(type) {
	case int:
		return f.(int) * 2
	case string:
		return f.(string) + f.(string) + f.(string) + f.(string)
	}
	return f
}
