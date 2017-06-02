package solver

import (
	"github.com/step17/ihazwordz/words"
)

type PointMap map[rune]int

var LetterPoints PointMap

func init() {
	LetterPoints = make(PointMap)
	// Assign everything 1 for now:
	for c := 'A'; c <= 'Z'; c++ {
		LetterPoints[c] = 1
	}
	// Assign some letters more.
	LetterPoints.assignPoints("LCFHMPVWY", 2)
	LetterPoints.assignPoints("JKQXZ", 3)
}

func (m PointMap) assignPoints(letters string, points int) {
	for _, l := range letters {
		m[l] = points
	}
}

func Score(word string) int {
	if word == "" {
		return 0
	}
	score := 1
	for _, l := range words.Normalize(word) {
		score += LetterPoints[l]
	}
	return score * score
}
