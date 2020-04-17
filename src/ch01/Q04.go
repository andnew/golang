package ch01

import (
	"fmt"
)

func Q4() {
	fmt.Println("=========Q04=========")
	avg := t41([]float64{4.0, 1.0, 2.0, 3.0})
	fmt.Println(avg)
}

func t41(xs []float64) float64 {

	sum := 0.0
	avg := 0.0

	switch len(xs) {
	case 0:
		avg = 0.0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	return avg
}
