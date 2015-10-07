package main

import (
	"fmt"
	"log"
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

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	numWords, wordlist := usage()

	g, err := diceware.GeneratorFromFile(wordlist)
	fatalIfErr(err)

	n, err := strconv.Atoi(numWords)
	fatalIfErr(err)

	dwp, err := g.Generate(n)
	fatalIfErr(err)

	fmt.Println(dwp)
}
