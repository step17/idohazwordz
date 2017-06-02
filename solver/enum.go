package solver

import "github.com/step17/ihazwordz/words"

// EnumSolver enumerates all of the subsets of a given string
// iteratively instead of recursively.
type EnumSolver struct {
	rs RecursiveSolver // Borrowing RecursiveSolver's choice list.
	Solver
}

func (s *EnumSolver) Init(dict []string) error {
	return s.rs.Init(dict)
}

func (s *EnumSolver) Solve(letters string) string {
	sorted := words.Sort(letters)
	max := 1 << uint(len(letters))
	var best *choices
	for i := 0; i < max; i++ {
		var sub string
		for j := uint(0); j < uint(len(letters)); j++ {
			if (i & (1 << j)) != 0 {
				sub += string(sorted[j])
			}
		}
		cs := s.rs.sorted[sub]
		if best.score() < cs.score() {
			best = cs
		}
	}
	return best.first()
}
