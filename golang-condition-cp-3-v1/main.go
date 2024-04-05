package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {

	ratarata := (math + science + english + indonesia) / 4

	switch {
	case ratarata == 100:
		return "Sempurna"
	case ratarata >= 90:
		return "Sangat Baik"
	case ratarata >= 80:
		return "Baik"
	case ratarata >= 70:
		return "Cukup"
	case ratarata >= 60:
		return "Kurang"
	default:
		return "Sangat kurang"
	}
	// TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
