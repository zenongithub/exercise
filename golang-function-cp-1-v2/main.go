package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {
	months := [12]string{"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}

	dayStr := fmt.Sprintf("%02d", day)

	monthName := months[month-1]

	formattedDate := fmt.Sprintf("%s-%s-%d", dayStr, monthName, year)

	return formattedDate
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
