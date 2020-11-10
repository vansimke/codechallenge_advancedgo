package main

import (
	"challenge/book"
	"challenge/stats"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	option := flag.Int("option", 1, "1 - read all text, 2 - read first 20 lines, 3 - analyze word frequency")
	flag.Parse()

	f, err := os.Open("./mobydick.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	b := book.New(f, stats.Stats{})
	switch *option {
	case 1:
		fmt.Println(b.Read())
	case 2:
		lines, err := b.ReadLines(0, 20)
		if err != nil {
			log.Fatal(err)
		}
		for _, l := range lines {
			fmt.Println(l)
		}
	case 3:
		fmt.Println(b.WordFrequency())
	}
}
