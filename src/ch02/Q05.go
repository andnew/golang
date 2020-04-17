package ch02

import "fmt"

//Q5. (0) 平均值
//1. 编写一个函数用于计算一个 float64 类型的 slice 的平均值。

func Q5() {
	fmt.Println("=========Q05==========")
	t51()
}

func t51() {

	a := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	sum := 0.0

	for _, v := range a {
		sum += v
	}

	avg := sum / float64(len(a))

	fmt.Println(avg)
}
