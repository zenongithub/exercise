package main

import "fmt"

func GraduateStudent(score int, absent int) string {
	if score >= 70 && absent < 5 {
		return "lulus"
	} else if score < 70 || absent >= 5 {
		return "tidak lulus"
	}
	return "" // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GraduateStudent(80, 5))
}
