package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {
	switch v := time.(type) {
	case string:
		hour, minute, err := parseStringTime(v)
		if err != nil {
			return "Invalid input"
		}
		return convertTo12HourFormat(hour, minute)
	case []int:
		if len(v) != 2 {
			return "Invalid input"
		}
		return convertTo12HourFormat(v[0], v[1])
	case map[string]int:
		hour, ok1 := v["hour"]
		minute, ok2 := v["minute"]
		if !ok1 || !ok2 {
			return "Invalid input"
		}
		return convertTo12HourFormat(hour, minute)
	case Time:
		return convertTo12HourFormat(v.Hour, v.Minute)
	default:
		return "Invalid input"
	}
}

func parseStringTime(timeStr string) (int, int, error) {
	timeComponents := [...]string{"Hour", "Minute"}
	timeValues := make(map[string]int)

	splitTime := strings.Split(timeStr, ":")
	if len(splitTime) != 2 {
		return 0, 0, fmt.Errorf("invalid time format")
	}

	for i, component := range splitTime {
		value, err := strconv.Atoi(component)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid time format")
		}
		timeValues[timeComponents[i]] = value
	}

	return timeValues["Hour"], timeValues["Minute"], nil
}

func convertTo12HourFormat(hour, minute int) string {
	if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return "Invalid input"
	}

	period := "AM"
	if hour >= 12 {
		period = "PM"
	}

	if hour == 0 {
		hour = 12
	} else if hour > 12 {
		hour -= 12
	}

	return fmt.Sprintf("%02d:%02d %s", hour, minute, period)
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}
