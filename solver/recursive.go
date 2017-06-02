package solver

import (
	"math"

	"github.com/step17/ihazwordz/words"
)

// Choices is a type for keeping track of lists of word choices with
// equivalent scores.
type choices struct {
	words  []string
	points int
}

// Score returns the score of this choice.
func (cs *choices) score() int {
	if cs == nil {
		return 0
	}
	return cs.points
}

// First arbitrarily picks the first out of a list of potential word
// choices, or an empty string if there are no choices.
func (cs *choices) first() string {
	if cs == nil {
		return ""
	}
	return cs.words[0]
}

// Basic recursive solver. Tries all 2^N possible subsets of letters.
type RecursiveSolver struct {
	sorted map[string]*choices
	minLen int
	Solver
}

func (s *RecursiveSolver) Init(dict []string) error {
	s.sorted = make(map[string]*choices)
	s.minLen = math.MaxInt32
	for _, word := range dict {
		if len(word) < s.minLen {
			s.minLen = len(word)
		}
		norm := words.Normalize(word)
		v := s.sorted[words.Sort(norm)]
		if v == nil {
			v = &choices{points: Score(norm)}
			s.sorted[words.Sort(norm)] = v
		}
		v.words = append(v.words, norm)
	}
	return nil
}

func (s *RecursiveSolver) Solve(letters string) string {
	cs := s.resolve("", words.Sort(letters))
	return cs.first()
}

func (s RecursiveSolver) resolve(picked, remain string) *choices {
	if remain == "" {
		return s.sorted[picked]
	}
	if len(picked)+len(remain) < s.minLen {
		return nil
	}
	next := remain[1:]
	skip := s.resolve(picked, next)
	kept := s.resolve(picked+remain[:1], next)
	if skip.score() > kept.score() {
		return skip
	}
	return kept
}
