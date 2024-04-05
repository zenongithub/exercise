package main

import "fmt"

func BMICalculator(gender string, height int) float64 {

	var persen float64

	if gender == "laki-laki" {
		persen = 0.10
	} else if gender == "perempuan" {
		persen = 0.15
	} else {
		fmt.Println("Gender tidak valid")
		return 0.0
	}

	hasil := float64(height-100) - float64(height-100)*persen

	return hasil
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 165))
	fmt.Println(BMICalculator("perempuan", 165))
}
