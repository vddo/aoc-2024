package solver

type GridSolver interface {
	Solve() error
	Render()
}
