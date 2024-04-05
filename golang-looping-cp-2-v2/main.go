package main

import "fmt"

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {

	var reversed string

	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	var final string

	for i := 0; i < len(reversed); i++ {
		if (i+1 < len(reversed) && reversed[i+1] == ' ') || reversed[i] == ' ' {
			final += string(reversed[i])
		} else {
			final += string(reversed[i]) + "_"
		}
	}

	return final[:len(final)-1]
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
