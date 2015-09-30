package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/garrettr/diceware"
)

func usage() (num_words, wordlist string) {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <numWords> <wordlist>\n", os.Args[0])
		os.Exit(1)
	}
	return os.Args[1], os.Args[2]
}

func main() {
	numWords, wordlist := usage()

	g, err := diceware.GeneratorFromFile(wordlist)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	n, err := strconv.Atoi(numWords)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dwp, err := g.Generate(n)
	fmt.Println(dwp)
}
