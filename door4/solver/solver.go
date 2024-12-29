package solver

import (
	"door4/arraystack"
	"strings"
)

const (
	KEYWORD string = "XMAS"
)

var bound struct {
	rows    int
	columns int
}

// Checks if the argument is the first letter of the KEYWORD.
// Sensitive to capital letters.
func beginning(letter byte) bool {
	return letter == KEYWORD[0]
}

// Returns next letter that follows argument in KEYWORD
func nextLetter(letter byte) byte {
	i := strings.IndexByte(KEYWORD, letter)
	if i < len(KEYWORD)-1 {
		return KEYWORD[i+1]
	}
	return 0
}

// Check if indices are in bound of data
func inBound(m int, n int) bool {
	return m >= 0 && m < bound.rows && n >= 0 && n < bound.columns
}

func Solve(data *[]string) (int, error) {
	stack, result := arraystack.ArrayStack{}, arraystack.ArrayStack{}

	countKeyword := 0
	bound.rows, bound.columns = len(*data), len((*data)[0])
	for i := 0; i < bound.rows; i++ {
		for j := 0; j < bound.columns; j++ {
			value := (*data)[i][j]
			if beginning(value) {
				stack.Push(&arraystack.Node{Val: value, I: i, J: j})
			}

			for !stack.Empty() {
				node := *(stack.Pop())
				next := nextLetter(node.Val)
				node_i, node_j := node.I, node.J

				if next == 0 {
					result.Push(&node)
					countKeyword++
					continue
				}

				if node.Direction != nil {
					m, n := node_i+node.Direction.V_x, node_j+node.Direction.V_y
					if !inBound(m, n) {
						continue
					}

					if neighbor := (*data)[m][n]; neighbor == next {
						stack.Push(&arraystack.Node{
							Val:       neighbor,
							I:         m,
							J:         n,
							Parent:    &node,
							Direction: node.Direction,
						})
					}
					continue
				}

				for g := 0; g < 3; g++ {
					for h := 0; h < 3; h++ {
						m, n := node_i-1+g, node_j-1+h
						if !inBound(m, n) {
							continue
						}

						if neighbor := (*data)[m][n]; neighbor == next {
							stack.Push(&arraystack.Node{
								Val:       neighbor,
								I:         m,
								J:         n,
								Parent:    &node,
								Direction: &arraystack.Vec{V_x: m - node_i, V_y: n - node_j},
							})
						}
					}
				}
			}
		}
	}

	return countKeyword, nil
}
