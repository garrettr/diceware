package diceware

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func strSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestNewFromSlice(t *testing.T) {
	testSlice := []string{"a", "b", "c"}
	g := GeneratorFromSlice(testSlice)
	if !strSlicesEqual(testSlice, g.words) {
		t.Errorf("GeneratorFromSlice should have words=%v, found=%v", testSlice, g.words)
	}
}

func TestNewFromFile(t *testing.T) {
	testWords := []string{"a", "b", "c"}

	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Failed to create temporary file with error: %s", err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())

	tf.WriteString(strings.Join(testWords, "\n"))
	g, err := GeneratorFromFile(tf.Name())
	if err != nil {
		t.Fatalf("Failed to create GeneratorFromFile with error: %s", err)
	}

	// The generator's word list should match the slice of words we wrote to the file
	if !strSlicesEqual(testWords, g.words) {
		t.Errorf("GeneratorFromFile should have words=%v, found=%v", testWords, g.words)
	}
}

func strSliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TestGenerate(t *testing.T) {
	testWords := []string{"a", "b", "c"}
	testPhraseLength := 4
	g := GeneratorFromSlice(testWords)
	dw, err := g.Generate(testPhraseLength)
	if err != nil {
		t.Error(err)
	}

	// The generated Diceware passphrase should have the expected length
	if (len(strings.Split(dw, " "))) != testPhraseLength {
		t.Errorf("Generated Diceware passphrase should have length %d, found '%s' with length %d", testPhraseLength, dw, len(strings.Split(dw, " ")))
	}

	// Each element of the generated passphrase should be from the
	// slice of strings used to create the generator
	for _, word := range strings.Split(dw, " ") {
		if !strSliceContains(testWords, word) {
			t.Errorf("Found unexpected word '%s' in generated phrase '%s'", word, dw)
		}
	}
}
