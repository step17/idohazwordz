package solver

type Solver interface {
	Init(dict []string) error
	Solve(letters string) string
}

var kAllSolvers = []Solver{
	&RecursiveSolver{},
	&MemoSolver{},
	&EnumSolver{},
	&ListSolver{},
	&CountListSolver{},
	&BitfieldSolver{},
}
