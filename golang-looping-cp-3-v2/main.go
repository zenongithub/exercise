package main

import (
	"fmt"
	"strings"
)

func CountingLetter(text string) int {

	invalidLetters := "RSTZrstz"

	count := 0
	for _, char := range text {
		if strings.ContainsRune(invalidLetters, char) {
			count++
		}
	}
	return count

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
