package solver

import (
	"fmt"
	"testing"

	"github.com/step17/ihazwordz/words"
)

func testSolver(t *testing.T, s Solver) {
	err := s.Init([]string{"cat", "dog", "vim", "pomplamoose", "queen"})
	if err != nil {
		t.Fatalf("Init() err'd: %v", err)
	}
	for _, tc := range []struct {
		in   string
		want string // something scoring at least as many points as this
	}{
		{"CAT", "CAT"},
		{"ACT", "CAT"},
		{"ACTQ", "CAT"},
		{"DFQ", ""},
		{"", ""},
		{"VIMCAT", "VIM"},
		{"MCIVTA", "VIM"},
		{"OMPPALSEOOMQK", "POMPLAMOOSE"},
		{"NEEQ", "QUEEN"},
	} {
		got := s.Solve(tc.in)
		if !words.Count(tc.in).Contains(words.Count(got)) {
			t.Errorf("Solve(%q) = %q but %q does not contain %q", tc.in, got, tc.in, got)
			continue
		}
		gotScore := Score(got)
		wantScore := Score(tc.want)
		if gotScore < wantScore {
			t.Errorf("Solve(%q) = %q (score=%v). Wanted at least %q (score=%v)", tc.in, got, gotScore, tc.want, wantScore)
		}
	}
}

func TestAll(t *testing.T) {
	for _, s := range kAllSolvers {
		t.Run(fmt.Sprintf("%T", s), func(t *testing.T) { testSolver(t, s) })
	}
}
