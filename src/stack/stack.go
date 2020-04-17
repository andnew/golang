package stack

type Stack struct {
	i    int   //长度
	data []int // 数据容器
}

func (s *Stack) Push(v int) {
	if s.data == nil {
		s.data = make([]int, 0)
		s.i = 0
	}
	s1 := make([]int, s.i, s.i)
	copy(s1, s.data)
	s.data = s1
	s.data = append(s.data, v)
	s.i = s.i + 1

}

func (s *Stack) Pop() (res int) {
	if s.i == 0 {
		return -1
	}
	res = s.data[s.i-1]
	s.i = s.i - 1
	return
}
