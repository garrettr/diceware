package diceware

import (
	"bufio"
	"crypto/rand"
	"math/big"
	"os"
	"strings"
)

type DicewareGenerator struct {
	words []string
}

func GeneratorFromSlice(slice []string) *DicewareGenerator {
	return &DicewareGenerator{slice}
}

func GeneratorFromFile(fname string) (*DicewareGenerator, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return GeneratorFromSlice(lines), nil
}

// randInt Wraps crypto/rand's rand.Int so it's easier to use with
// ints, like those that we get when working with the lengths and
// indexes of slices.
func randInt(max int) (int, error) {
	maxBig := *big.NewInt(int64(max))
	n, err := rand.Int(rand.Reader, &maxBig)
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}

func (d DicewareGenerator) Generate(words int) (string, error) {
	var randWords []string
	for i := 0; i < words; i++ {
		ri, err := randInt(len(d.words))
		if err != nil {
			return "", err
		}
		randWords = append(randWords, d.words[ri])
	}
	return strings.Join(randWords, " "), nil
}
