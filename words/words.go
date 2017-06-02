package words

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type CountMap map[string]int

var (
	qFixer   = strings.NewReplacer("QU", "Q")
	qUnfixer = strings.NewReplacer("Q", "Qu")
	wordRE   = regexp.MustCompile(`^(?i:(?:[a-pr-z])|qu){3,}$`)
)

const (
	kMinLen = 3
)

func Normalize(word string) string {
	return qFixer.Replace(strings.ToUpper(word))
}

func Denormalize(word string) string {
	return qUnfixer.Replace(Normalize(word))
}

func Count(word string) CountMap {
	return CountLetters(strings.Split(Normalize(word), ""))
}

func CountLetters(letters []string) CountMap {
	c := make(CountMap)
	for _, l := range letters {
		c[l]++
	}
	return c
}

func LoadValidFile(fileName string, maxLen int, ch chan string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	LoadValid(file, maxLen, ch)
	return nil
}

func LoadValid(reader io.Reader, maxLen int, ch chan string) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		word := scanner.Text()
		if !wordRE.MatchString(word) {
			continue
		}
		norm := Normalize(word)
		if len(norm) > maxLen || len(norm) < kMinLen {
			continue
		}
		ch <- word
	}
}

func Load(filename string, maxLen int) []string {
	var dict []string
	ch := make(chan string, 32)
	go func() {
		LoadValidFile(filename, maxLen, ch)
		close(ch)
	}()
	for word := range ch {
		dict = append(dict, word)
	}
	log.Printf("loaded %v words", len(dict))
	return dict
}

// Finds the maximum of each letter count for a set of CountMaps.  In
// other words, the result of Max(a,b,c) would necessarily Contain any
// of a b and c for any a b c.
func Max(counts ...CountMap) CountMap {
	res := make(CountMap)
	for _, cm := range counts {
		for l, c := range cm {
			if res[l] > c {
				continue
			}
			res[l] = c
		}
	}
	return res
}

func (haystack CountMap) Contains(needle CountMap) bool {
	for l, c := range needle {
		if haystack[l] < c {
			return false
		}
	}
	return true
}

func (cm CountMap) String() string {
	var letters []string
	for l, c := range cm {
		for i := 0; i < c; i++ {
			letters = append(letters, l)
		}
	}
	sort.Sort(sort.StringSlice(letters))
	return strings.Join(letters, "")
}
