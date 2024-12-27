package arraystack

import "errors"

type ArrayStack []rune

func New() *ArrayStack {
	return &ArrayStack{}
}

func (s *ArrayStack) Push(x rune) {
	*s = append(*s, x)
}

func (s *ArrayStack) Pop() (rune, error) {
	length := len(*s)
	if length == 0 {
		return 0, errors.New("empty stack")
	}

	r := (*s)[length-1]
	*s = (*s)[:length-1]

	return r, nil
}

func (s *ArrayStack) Empty() bool {
	return len(*s) == 0
}
