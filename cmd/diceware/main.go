package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/garrettr/diceware"
)

func parseArgs() (int, string) {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <numWords> <wordlist>\n", os.Args[0])
		os.Exit(1)
	}

	numWords, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid value for numWords: '%s'", os.Args[1])
	}

	return numWords, os.Args[2]
}

func main() {
	numWords, wordlist := parseArgs()

	g, err := diceware.FromFile(wordlist)
	if err != nil {
		log.Fatalf("Could not create generator from wordlist '%s': %s", wordlist, err)
	}

	dw, err := g.Generate(numWords)
	if err != nil {
		log.Fatalf("Error generating Diceware passphrase: %s", err)
	}

	fmt.Println(dw)
}
