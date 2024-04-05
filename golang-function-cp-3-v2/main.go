package main

import (
	"fmt"
	"strings"
)

func FindShortestName(names string) string {

	var shortestName string
	shortestLength := -1

	nameList := strings.FieldsFunc(names, func(r rune) bool {
		return r == ' ' || r == ',' || r == ';'
	})

	for _, name := range nameList {
		length := len(name)
		if shortestLength == -1 || length < shortestLength || (length == shortestLength && name < shortestName) {
			shortestName = name
			shortestLength = length
		}
	}

	return shortestName
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
