package ch02

import "fmt"

func Q11() {
	fmt.Println("=========Q11=========")
	t111()
}

func t111() {
	m := []int{1, 3, 4}
	f := func(i int) int {
		return i * i
	}
	fmt.Printf("%v\n", Map(f, m))
}

func Map(f func(int) int, l []int) []int {
	j := make([]int, len(l))
	for k, v := range l {
		j[k] = f(v)
	}
	return j
}
