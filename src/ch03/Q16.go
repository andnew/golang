package ch03

import (
	"bufio"
	"fmt"
	"os"
	"stack"
	"strconv"
)

//使用 stack 包创建逆波兰计算器。

func Q16() {
	t161()
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var st = stack.Stack{}

func t161() {
	for {
		s, err := reader.ReadString('\n')
		var token string
		if err != nil {
			return
		}
		fmt.Println("输入字符 ", s)
		for _, c := range s {
			switch {
			case c >= '0' && c <= '9':
				token = token + string(c)
			case c == ' ':
				r, _ := strconv.Atoi(token)
				st.Push(r)
				token = ""
			case c == '+':
				fmt.Printf("%d\n", st.Pop()+
					st.Pop())
			case c == '*':
				fmt.Printf("%d\n", st.Pop()*
					st.Pop())
			case c == '-':
				p := st.Pop()
				q := st.Pop()
				fmt.Printf("%d\n", q-p)
			case c == 'q':
				return
			default:
				//error
			}
		}
	}
}
