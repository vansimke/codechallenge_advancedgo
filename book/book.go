package book

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type WordFreqCalc interface {
	WordFrequency([]string) map[string]int
}

type Book struct {
	contents     []string
	wordFreqCalc WordFreqCalc
}

func (b Book) Read() string {
	return strings.Join(b.contents, "\n")
}

// ReadLines returns 'num' lines from the book, starting at line 'start' (base 0).
// uints are used for parameters to eliminate the need to check for negative numbers
func (b Book) ReadLines(start, num uint) ([]string, error) {
	if start >= uint(len(b.contents)) {
		return nil, errors.New("Starting value greater than lines in book")
	}
	if start+num > uint(len(b.contents)) {
		num = uint(len(b.contents)) - start
	}
	// return a copy of the contents to ensure that consumer doesn't manipulate contents
	result := make([]string, num)
	copy(result, b.contents[start:start+num])
	return result, nil
}

func (b Book) WordFrequency() map[string]int {
	return b.wordFreqCalc.WordFrequency(b.contents)
}

func New(source io.Reader, wordFreqCalc WordFreqCalc) *Book {
	s := bufio.NewScanner(source)

	// give the contents slice a reasonable starting size to limit reallocations
	b := Book{
		contents:     make([]string, 0, 5000),
		wordFreqCalc: wordFreqCalc,
	}
	for s.Scan() {
		b.contents = append(b.contents, s.Text())
	}

	return &b
}
