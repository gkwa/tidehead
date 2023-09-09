package main

import (
	"fmt"
	"time"
)

func main() {
	i := int64(100000)
	e := time.Unix(i, 0)
	now := time.Now()
	d1 := now.Sub(e)

	fmt.Printf("%v, %s\n", d1, formatDuration(d1))

	d1 = 10 * time.Minute
	fmt.Printf("%v: %s\n", d1, formatDuration(d1))
	d1 = 10*time.Minute + 10*time.Second
	fmt.Printf("%v: %s\n", d1, formatDuration(d1))
	d1 = 10*time.Minute + 61*time.Second
	fmt.Printf("%v: %s\n", d1, formatDuration(d1))

	tenYearsInSeconds := int64(10 * 365 * 24 * 60 * 60)
	d1 = time.Duration(tenYearsInSeconds) * time.Second
	fmt.Printf("%v: %s\n", d1, formatDuration(d1))

	d1 = time.Duration(int64(1*24*60*60)) * time.Second
	fmt.Printf("%v: %s\n", d1, formatDuration(d1))
}

func formatDuration(duration time.Duration) string {
	days := int(duration.Hours()) / 24
	months := days / 30
	days %= 30
	years := months / 12
	months %= 12
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	result := ""

	if years > 0 {
		result += fmt.Sprintf("%dy", years)
	}
	if months > 0 {
		result += fmt.Sprintf("%dm", months)
	}
	if days > 0 {
		result += fmt.Sprintf("%dd", days)
	}
	if hours > 0 {
		result += fmt.Sprintf("%dh", hours)
	}
	if minutes > 0 || result == "" {
		result += fmt.Sprintf("%dm", minutes)
	}

	return result

}
