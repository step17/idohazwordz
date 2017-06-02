package solver

import (
	"math/big"
	"sort"

	"github.com/step17/ihazwordz/words"
)

// Bitfield here describes a bitfield-style representation of the
// number of each letter required to spell a word, or are available on
// a board. For example, if our dictionary consisted of three words:
// "A", "AA", and "AB" if we had a board of letters "AAB" we could
// spell anyword in the dictionary, and we'll call that the "max
// covering". The bitfield expression of the "max covering" is all 1s,
// so: "AAB" -> "111". Repeated letters are filled from
// least-significant to most-significant digits, so "A" is expressed
// as "010". Whether one word can be spell from another can then be
// expressed as "board & word == word" e.g. "AA" ("110") can spell both "A"
// ("010") and "AA" ("110") but not "AB" ("011").
type bitfield struct {
	i *big.Int
}

// Contains returns true iff needle is a subset of haystack.
func (haystack *bitfield) Contains(needle bitfield) bool {
	var i big.Int
	return needle.i.Cmp(i.And(haystack.i, needle.i)) == 0
}

// bfPair stores an anagram word pair (the original word, and a count
// of the required letters expressed as a bitfield).
type bfPair struct {
	bf   bitfield
	word string
}

// bfPairs sortable by score.
type bfPairSlice []bfPair

func (p bfPairSlice) Len() int      { return len(p) }
func (p bfPairSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p bfPairSlice) Less(i, j int) bool {
	return Score(p[j].word) < Score(p[i].word)
}

// BitfieldSolver goes through a sorted dictionary like ListSolver,
// but uses bitfields instead of sorted strings to represent anagram
// clusters.
type BitfieldSolver struct {
	rs   RecursiveSolver
	dict []bfPair       // Highest scoring words first.
	max  words.CountMap // Max-covering count for the whole dictionary
	Solver
}

func (s *BitfieldSolver) Init(dict []string) error {
	s.rs.Init(dict)
	for sorted := range s.rs.sorted {
		s.max = words.Max(s.max, words.Count(sorted))
	}
	for sorted, cs := range s.rs.sorted {
		// all of thse are equivalently high scoring anagrams of
		// eachother, so just pick one.
		s.dict = append(s.dict, bfPair{bf: s.Bitfield(sorted), word: cs.first()})
	}
	sort.Sort(bfPairSlice(s.dict))
	return nil
}

// Expresses the given word as a bitfield.
func (s BitfieldSolver) Bitfield(word string) bitfield {
	cm := words.Count(word)
	i := big.NewInt(0)
	for li := 'A'; li <= 'Z'; li++ {
		l := string(li)
		width := uint(s.max[l])
		if width == 0 {
			// l is not used in dictionary
			continue
		}
		fill := (1 << uint(cm[l])) - 1
		i.Lsh(i, width)
		i.Or(i, big.NewInt(int64(fill)))
	}
	return bitfield{i}
}

func (s BitfieldSolver) Solve(letters string) string {
	bf := s.Bitfield(letters)
	for _, p := range s.dict {
		if bf.Contains(p.bf) {
			return p.word
		}
	}
	return ""
}
