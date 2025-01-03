package arraystack

type ArrayStack []*Node

type Node struct {
	Val       byte
	X, Y      int
	Parent    *Node
	Direction *Vec
}

type Vec struct {
	X, Y int
}

func New() *ArrayStack {
	return &ArrayStack{}
}

func (as *ArrayStack) Push(x *Node) {
	*as = append(*as, x)
}

func (as *ArrayStack) Pop() *Node {
	length := len(*as)
	if length == 0 {
		return nil
	}

	n := (*as)[length-1]
	*as = (*as)[:length-1]

	return n
}

func (as *ArrayStack) Empty() bool {
	return len(*as) == 0
}
