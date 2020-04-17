package ch02

import "fmt"

//Q5. (0) 平均值
//1. 编写一个函数用于计算一个 float64 类型的 slice 的平均值。

func Q5() {
	fmt.Println("=========Q05==========")
	a := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	t51(a)
	avg := t52(a)
	fmt.Println("切片 ", a, " , avg = ", avg)
}

func t51(nums []float64) {

	sum := 0.0
	for _, v := range nums {
		sum += v
	}

	avg := sum / float64(len(nums))

	fmt.Println(avg)
}

func t52(nums []float64) (avg float64) {
	// 官方的答案
	fmt.Println("官方的答案")

	switch len(nums) {
	case 0:
		avg = 0.0
	default:
		sum := 0.0
		for _, v := range nums {
			sum += v
		}
		avg = sum / float64(len(nums))
	}
	return
}
