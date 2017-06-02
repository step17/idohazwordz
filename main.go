package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/step17/ihazwordz/solver"
	"github.com/step17/ihazwordz/words"
)

var (
	dictFile = flag.String("dictionary", "/usr/share/dict/words", "File to use as a dictionary")
	size     = flag.Int("size", 16, "How many letters to expect")
)

func main() {
	s := &solver.ListSolver{}
	s.Init(words.Load(*dictFile, *size))

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Input letters: ")
		input, err := reader.ReadString('\n')
		if input == "\n" || err != nil {
			break
		}
		fmt.Printf("Best answer: %s\n", s.Solve(words.Normalize(input)))
	}
}
