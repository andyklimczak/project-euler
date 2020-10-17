package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC)
	end:= time.Date(2000, time.December, 30, 0, 0, 0, 0, time.UTC)
	count := 0
	for date := start; date.Before(end); date = date.AddDate(0, 1, 0) {
		if date.Weekday() == time.Sunday {
			count++
		}
	}
	fmt.Println(count)
}
