package stack

import "testing"

func TestStack_Push(t *testing.T) {
	stack := Stack{}
	if -1 == stack.Pop() {
		t.Log("空栈返回的数据值 ", stack.Pop())
	}
	stack.Push(1)
	if 1 == stack.i {
		t.Log("栈的高度 ", stack.Pop())
	}

	if 1 == stack.Pop() {
		t.Log("入栈后，栈第一次pop值是 1")
	}
}

func TestStackPop(t *testing.T) {

	s := new(Stack)
	s.Push(3)
	s.Pop()
	s.Push(4)

	t.Log("栈的元素 ", s.data)
	if 1 != s.i {
		t.Log("栈的高度计算错误")
		t.Fail()
	}

}
