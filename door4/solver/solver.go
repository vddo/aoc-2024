package solver

import (
	"door4/arraystack"
	"fmt"
	"strings"
	"sync"
)

const (
	KEYWORD string = "XMAS"
	SPACE   byte   = 32
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

type DirectionCache struct {
	Directions []*arraystack.Vec
	Once       sync.Once
}

var GlobalDirectionCache *DirectionCache = &DirectionCache{}

func NewSolver(data *[]string, keyword string, rows int, columns int) *Solver {
	return &Solver{Data: data, Keyword: keyword, Rows: rows, Columns: columns}
}

func (dc *DirectionCache) GetDirections() {
	dc.Once.Do(func() {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i != 0 || j != 0 {
					dc.Directions = append(dc.Directions, &arraystack.Vec{X: i, Y: j})
				}
			}
		}
	})
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
func (s *Solver) inBound(m int, n int) bool {
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
			grid[i][j] = SPACE
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

	*list = append(*list, GridPoint{node.X, node.Y, node.Val})
}

func printOneRow(row *[]byte) {
	for i := range *row {
		fmt.Printf("%c", (*row)[i])
	}
	fmt.Println("")
}

func (s *Solver) Solve() error {
	GlobalDirectionCache.GetDirections()

	for i := 0; i < s.Rows; i++ {
		for j := 0; j < s.Columns; j++ {
			value := (*s.Data)[i][j]
			if isFirstLetter(value) {
				s.SearchStack.Push(&arraystack.Node{Val: value, X: i, Y: j})
			}

			for !s.SearchStack.Empty() {
				node := *(s.SearchStack.Pop())
				next := nextLetter(node.Val)

				if next == 0 {
					getPath(&node, &s.Result)
					s.KeywordCount++
					continue
				}

				if node.Direction != nil {
					m, n := node.X+node.Direction.X, node.Y+node.Direction.Y
					if !s.inBound(m, n) {
						continue
					}

					if neighbor := (*s.Data)[m][n]; neighbor == next {
						s.SearchStack.Push(&arraystack.Node{
							Val:       neighbor,
							X:         m,
							Y:         n,
							Parent:    &node,
							Direction: node.Direction,
						})
					}
					continue
				}

				for _, d := range GlobalDirectionCache.Directions {
					m, n := node.X+d.X, node.Y+d.Y
					if !s.inBound(m, n) {
						continue
					}

					if neighbor := (*s.Data)[m][n]; neighbor == next {
						s.SearchStack.Push(&arraystack.Node{
							Val:       neighbor,
							X:         m,
							Y:         n,
							Parent:    &node,
							Direction: &arraystack.Vec{X: m - node.X, Y: n - node.Y},
						})
					}
				}

			}
		}
	}

	return nil
}
