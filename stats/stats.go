package stats

import (
	"log"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Stats struct{}

func (s Stats) TotalLines(lines []string) int {
	return len(lines)
}

func (s Stats) TotalWords(lines []string) int {
	count := 0
	for _, line := range lines {
		count += len(strings.Fields(line))
	}
	return count
}

func (s Stats) WordFrequency(lines []string) map[string]int {
	lineCh := make(chan string)
	m := sync.Mutex{}
	var result = make(map[string]int)

	g := new(errgroup.Group)

	for i := 0; i < 10; i++ {
		g.Go(func() error {
			for line := range lineCh {
				for _, word := range strings.Fields(line) {
					m.Lock()
					result[word]++
					m.Unlock()
				}
			}
			return nil
		})
	}

	for _, line := range lines {
		lineCh <- line
	}
	close(lineCh)
	err := g.Wait()
	if err != nil {
		log.Fatal(err)
	}

	return result
}
