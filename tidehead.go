package tidehead

import (
	"fmt"
	"log/slog"
	"time"
)

func FormatDuration(d time.Duration) string {
	db := DurationBreakdown{Input: d}

	const daysInMonth = 30 // Approximate

	db.Years = d / (365 * 24 * time.Hour)
	d -= db.Years * 365 * 24 * time.Hour
	db.Months = d / (daysInMonth * 24 * time.Hour)
	d -= db.Months * daysInMonth * 24 * time.Hour
	db.Days = d / (24 * time.Hour)
	d -= db.Days * 24 * time.Hour
	db.Hours = d / time.Hour
	d -= db.Hours * time.Hour
	db.Minutes = d / time.Minute
	d -= db.Minutes * time.Minute
	db.Seconds = d / time.Second

	result := ""

	if db.Years > 0 {
		result += fmt.Sprintf("%dy", db.Years)
	}
	if db.Months > 0 {
		result += fmt.Sprintf("%dM", db.Months)
	}
	if db.Days > 0 {
		result += fmt.Sprintf("%dd", db.Days)
	}
	if db.Years == 0 {
		if db.Hours > 0 {
			result += fmt.Sprintf("%dh", db.Hours)
		}
		if db.Days == 0 {
			if db.Minutes > 0 || result == "" {
				result += fmt.Sprintf("%dm", db.Minutes)
			}
			if db.Hours == 0 {
				if db.Seconds > 0 || result == "" {
					result += fmt.Sprintf("%ds", db.Seconds)
				}
			}
		}
	}
	db.Output = result

	slog.Debug(db.String())
	return result
}
