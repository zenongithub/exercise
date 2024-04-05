package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {

	emailSplit := strings.Split(email, "@")

	emailInfo := strings.Split(emailSplit[1], ".")

	fmt.Println(emailInfo)
	var TLD string

	if len(emailInfo) > 2 {
		TLD = strings.Join(emailInfo[1:], ".")
	} else {
		TLD = emailInfo[1]
	}

	fmt.Println(TLD)

	return "Domain: " + emailInfo[0] + " dan TLD: " + TLD
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
