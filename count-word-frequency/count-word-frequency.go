package main

import (
	"log"
	"strings"
	"unicode"
)

func main() {
	str := "Go is good, Go is great"
    
	freqMap := countFrequency(str)

	for word, freq := range freqMap {
		log.Printf("%s %d\n", word, freq)
	}
}

func countFrequency(s string) map[string]int {
	var b strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		} else {
			b.WriteByte(' ')
		}
	}

	words := strings.Fields(b.String())
	mp := make(map[string]int, len(words))

	for _, w := range words {
		mp[w]++
	}

	return mp
}