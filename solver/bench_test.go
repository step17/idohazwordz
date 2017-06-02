package solver

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/step17/ihazwordz/words"
)

const (
	kLen  = 16
	kSeed = 0
)

type workload []string

func BenchmarkAll(b *testing.B) {
	dict := words.Load("/usr/share/dict/words", kLen)
	work := fakeWorkload(1 << 13)
	for _, s := range kAllSolvers {
		s.Init(dict)
		b.Run(fmt.Sprintf("%T", s), SolverBenchmark(b, s, work))
	}
}

func sampleString() string {
	var pool string
	vals := map[int]string{
		1: "abdeginorstu",
		2: "lcfhmpvwy",
		3: "jkqxz",
	}
	for p, s := range vals {
		for i := 0; i < p; i++ {
			pool += s
		}
	}
	return pool
}

func fakeWorkload(size int) workload {
	rng := rand.New(rand.NewSource(kSeed))
	pool := sampleString()
	work := make(workload, size)
	for i := 0; i < size; i++ {
		var board string
		for l := 0; l < kLen; l++ {
			board += string(pool[rng.Intn(len(pool))])
		}
		work[i] = board
	}
	return work
}

func SolverBenchmark(b *testing.B, s Solver, w workload) func(b *testing.B) {
	return func(b *testing.B) {
		var wi int
		for i := 0; i < b.N; i++ {
			if wi >= len(w) {
				wi = 0
			}
			s.Solve(w[wi])
		}
	}
}
