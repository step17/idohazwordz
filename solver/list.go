package solver

import (
	"sort"

	"github.com/step17/ihazwordz/words"
)

// anaPair stores an anagram word pair (the original word, and the
// letters when sorted).
type anaPair struct {
	sorted string
	word   string
}

// anaPairs sortable by score.
type anaPairSlice []anaPair

func (p anaPairSlice) Len() int      { return len(p) }
func (p anaPairSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p anaPairSlice) Less(i, j int) bool {
	return Score(p[j].word) < Score(p[i].word)
}

// ListSolver goes through a sorted dictionary, checking if each word
// can be made out of the given characters.
type ListSolver struct {
	rs   RecursiveSolver
	dict []anaPair // Highest scoring words first.
	Solver
}

func (s *ListSolver) Init(dict []string) error {
	s.rs.Init(dict)
	for sorted, cs := range s.rs.sorted {
		// all of thse are equivalently high scoring anagrams of
		// eachother, so just pick one.
		s.dict = append(s.dict, anaPair{sorted: sorted, word: cs.first()})
	}
	sort.Sort(anaPairSlice(s.dict))
	return nil
}

func (s ListSolver) Solve(letters string) string {
	sorted := words.Sort(letters)
	for _, p := range s.dict {
		if s.canSpell(p.sorted, sorted) {
			return p.word
		}
	}
	return ""
}

func (s ListSolver) canSpell(sub, sup string) bool {
	// precondition: both sub and sup are sorted.
	bi, pi := 0, 0
	// Check each character of the substring (bi) to see if it's a
	// subsequence of the superstring (pi).
	for bi < len(sub) && pi < len(sup) {
		switch {
		case sub[bi] == sup[pi]:
			bi++
			pi++
		case sub[bi] < sup[pi]:
			// sub contains a character that isn't in sup
			return false
		default:
			// sup contains some letters that aren't in sub
			pi++
		}
	}
	// Did we successfully make it to the end of sub? If so, it's
	// entirely a subsequence of sup.
	return bi == len(sub)
}
