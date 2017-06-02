package solver

import "testing"

func TestPointsComplete(t *testing.T) {
	for c := 'A'; c <= 'Z'; c++ {
		if got := LetterPoints[c]; got < 1 {
			t.Errorf("LetterPoints[%q] = %v, want > 0 points", c, got)
		}
	}
}

func TestScore(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want int
	}{
		{"CAT", 25},
		{"cat", 25},
		{"GRAMOPHONE", 196},
		{"QUACK", 100},
		{"QACK", 100},
		{"", 0},
	} {
		if got := Score(tc.in); got != tc.want {
			t.Errorf("Score(%q) = %v, want %v", tc.in, got, tc.want)
		}
	}
}
