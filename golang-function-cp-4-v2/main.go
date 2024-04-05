package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {

	var similarData []string

	for _, item := range data {
		if strings.Contains(item, input) {
			similarData = append(similarData, item)
		}
	}

	return strings.Join(similarData, ",")
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
