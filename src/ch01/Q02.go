package ch01

import "fmt"

//Q2. (0) FizzBuzz
//1. 解决这个叫做 Fizz-Buzz[23] 的问题：
//编写一个程序，打印从 1 到 100 的数字。当是三个倍数就打印 “Fizz” 代替数字，当是5的倍数就打印 “Buzz”。当数字同时是三5的倍数时，打印 “FizzBuzz”

func Q2() {
	const (
		num3 = "Fizz"
		num5 = "Buzz"
	)

	for i := 1; i <= 100; i++ {
		num := ""
		if i%3 == 0 {
			num = num + num3
		}
		if i%5 == 0 {
			num = num + num5
		}
		if len(num) > 1 {
			fmt.Println(num)
		}
	}
}
