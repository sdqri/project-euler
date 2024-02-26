package main

import (
	"fmt"
	"time"
)

// end is inclusive
func CountWeekdayOnFirstOfMonth(start time.Time, end time.Time, weekday time.Weekday) int {
	n := 0
	for {
		if start.After(end) {
			break
		}
		if start.Weekday() == weekday && start.Day() == 1 {
			n++
		}
		start = start.Add(24 * time.Hour)
	}
	return n
}

func main() {
	startExecution := time.Now()
	start := time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, time.December, 31, 0, 0, 0, 0, time.UTC)
	n := CountWeekdayOnFirstOfMonth(start, end, time.Sunday)
	elapsed := time.Since(startExecution)

	fmt.Printf(
		"number of sundays from %v to %v = %v (elapse time = %v)\n",
		start.Format("2006-01-02"),
		end.Format("2006-01-02"),
		n,
		elapsed,
	)
}
