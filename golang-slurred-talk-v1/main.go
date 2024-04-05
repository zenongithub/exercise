package main

import "fmt"

func SlurredTalk(words *string) {

	var characters = *words

	var newWords string

	for idx := 0; idx < len(characters); idx = idx + 1 {

		var oneChar = string(characters[idx])

		if oneChar == "S" || oneChar == "R" || oneChar == "Z" {
			newWords += "L"
		} else if oneChar == "s" || oneChar == "r" || oneChar == "z" {
			newWords += "l"
		} else {
			newWords += oneChar
		}
	}
	*words = newWords
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Saya Zebra"
	SlurredTalk(&words)
	fmt.Println(words)
}
