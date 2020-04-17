package ch01

import (
	"fmt"
	"unicode/utf8"
)

/*
Q3. (1) 字符串
1. 建立一个 Go 程序打印下面的内容（到 100 个字符）：
A
AA
AAA
AAAA
AAAAA
AAAAAA
AAAAAAA
...
2. 建立一个程序统计字符串里的字符数量：
asSASA ddd dsjkdsjs dk
同时输出这个字符串的字节数。 提示： 看看 unicode/utf8 包。
3. 扩展/修改上一个问题的程序，替换位置 4 开始的三个字符为 “abc”。
4. 编写一个 Go 程序可以逆转字符串，例如 “foobar” 被打印成 “raboof”。 提示： 不
幸的是你需要知道一些关于转换的内容，参阅 “转换” 第 59 页的内容
*/
func Q3() {
	fmt.Println("=========Q03=========")
	t1()
	t2()
	t3()
	t4()
}

func t1() {

	a := ""
	for i := 1; i < 100; i++ {
		a = a + "A"
		fmt.Println(a)
	}
}

func t2() {
	str := "asSASA ddd dsjkdsjs dk"
	fmt.Println(utf8.RuneCount([]byte(str)))
}

func t3() {
	s := "asSASA ddd dsjkdsjs dk"
	r := []rune(s)
	copy(r[4:4+3], []rune("abc"))
	fmt.Println(string(r))
}

func t4() {
	s := "foobar"
	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	fmt.Println(string(r))
}
