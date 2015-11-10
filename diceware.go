package diceware

import (
	"bufio"
	"crypto/rand"
	"math/big"
	"os"
	"strings"
)

type Generator struct {
	words []string
}

func FromSlice(slice []string) *Generator {
	return &Generator{slice}
}

func FromFile(fname string) (*Generator, error) {
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

	return FromSlice(lines), nil
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

func (g Generator) Generate(length int) (string, error) {
	var words []string
	for i := 0; i < length; i++ {
		r, err := randInt(len(g.words))
		if err != nil {
			return "", err
		}
		words = append(words, g.words[r])
	}
	return strings.Join(words, " "), nil
}
