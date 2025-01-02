package solver

import (
	"door4/arraystack"
	"fmt"
	"strings"
)

const (
	KEYWORD string = "XMAS"
)

type GridPoint struct {
	X, Y  int
	Value byte
}

type Solver struct {
	Data                        *[]string
	SearchStack                 arraystack.ArrayStack
	Result                      []GridPoint
	Keyword                     string
	KeywordCount, Rows, Columns int
}

func NewSolver(data *[]string, keyword string, rows int, columns int) *Solver {
	return &Solver{Data: data, Keyword: keyword, Rows: rows, Columns: columns}
}

// Checks if the argument is the first letter of the KEYWORD.
// Sensitive to capital letters.
func isFirstLetter(letter byte) bool {
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

// Check if indices are in gridBound of data
func (s Solver) inBound(m int, n int) bool {
	return m >= 0 && m < s.Rows && n >= 0 && n < s.Columns
}

// Render KEYWORD and dots.
func (s *Solver) Render() {
	grid := make([][]byte, s.Rows)
	for i := range grid {
		grid[i] = make([]byte, s.Columns)
	}

	for i := 0; i < s.Rows; i++ {
		for j := 0; j < s.Columns; j++ {
			grid[i][j] = 32
		}
	}

	for _, point := range s.Result {
		grid[point.X][point.Y] = point.Value
	}

	for i := range grid {
		printOneRow(&grid[i])
	}
}

// Recursivly evaluate coordinates and value of parent node
func getPath(node *arraystack.Node, list *[]GridPoint) {
	if node.Parent != nil {
		getPath(node.Parent, list)
	}

	*list = append(*list, GridPoint{node.I, node.J, node.Val})
}

func printOneRow(row *[]byte) {
	for i := range *row {
		fmt.Printf("%c", (*row)[i])
	}
	fmt.Println("")
}

func (s *Solver) Solve() error {
	for i := 0; i < s.Rows; i++ {
		for j := 0; j < s.Columns; j++ {
			value := (*s.Data)[i][j]
			if isFirstLetter(value) {
				s.SearchStack.Push(&arraystack.Node{Val: value, I: i, J: j})
			}

			for !s.SearchStack.Empty() {
				node := *(s.SearchStack.Pop())
				next := nextLetter(node.Val)
				node_i, node_j := node.I, node.J

				if next == 0 {
					getPath(&node, &s.Result)
					s.KeywordCount++
					continue
				}

				if node.Direction != nil {
					m, n := node_i+node.Direction.V_x, node_j+node.Direction.V_y
					if !s.inBound(m, n) {
						continue
					}

					if neighbor := (*s.Data)[m][n]; neighbor == next {
						s.SearchStack.Push(&arraystack.Node{
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
						if !s.inBound(m, n) {
							continue
						}

						if neighbor := (*s.Data)[m][n]; neighbor == next {
							s.SearchStack.Push(&arraystack.Node{
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

	return nil
}
