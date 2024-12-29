package arraystack

type ArrayStack []*Node

type Node struct {
	Val       byte
	I         int
	J         int
	Parent    *Node
	Direction *Vec
}

type Vec struct {
	V_x int
	V_y int
}

func New() *ArrayStack {
	return &ArrayStack{}
}

func (s *ArrayStack) Push(x *Node) {
	*s = append(*s, x)
}

func (s *ArrayStack) Pop() *Node {
	length := len(*s)
	if length == 0 {
		return nil
	}

	n := (*s)[length-1]
	*s = (*s)[:length-1]

	return n
}

func (s *ArrayStack) Empty() bool {
	return len(*s) == 0
}
