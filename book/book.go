package book

import (
	"bufio"
	"challenge/stats"
	"errors"
	"io"
	"strings"
)

type Book struct {
	contents []string
	stats    stats.Stats
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
	return b.stats.WordFrequency(b.contents)
}

func New(source io.Reader, stats stats.Stats) *Book {
	s := bufio.NewScanner(source)

	// give the contents slice a reasonable starting size to limit reallocations
	b := Book{contents: make([]string, 0, 5000), stats: stats}
	for s.Scan() {
		b.contents = append(b.contents, s.Text())
	}
	return &b
}
