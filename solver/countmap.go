package solver

import (
	"sort"

	"github.com/step17/ihazwordz/words"
)

// cmPair stores an anagram word pair (the original word, and a CountMap of the letters).
type cmPair struct {
	cm   words.CountMap
	word string
}

// cmPairs sortable by score.
type cmPairSlice []cmPair

func (p cmPairSlice) Len() int      { return len(p) }
func (p cmPairSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p cmPairSlice) Less(i, j int) bool {
	return Score(p[j].word) < Score(p[i].word)
}

// CountListSolver goes through a sorted dictionary like CountListSolver,
// but uses words.countMaps instead of sorted strings to represent
// anagram clusters.
type CountListSolver struct {
	rs   RecursiveSolver
	dict []cmPair // Highest scoring words first.
	Solver
}

func (s *CountListSolver) Init(dict []string) error {
	s.rs.Init(dict)
	for sorted, cs := range s.rs.sorted {
		// all of thse are equivalently high scoring anagrams of
		// eachother, so just pick one.
		s.dict = append(s.dict, cmPair{cm: words.Count(sorted), word: cs.first()})
	}
	sort.Sort(cmPairSlice(s.dict))
	return nil
}

func (s CountListSolver) Solve(letters string) string {
	cm := words.Count(letters)
	for _, p := range s.dict {
		if cm.Contains(p.cm) {
			return p.word
		}
	}
	return ""
}
