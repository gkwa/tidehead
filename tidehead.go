package tidehead

import (
	"fmt"
	"log/slog"
	"time"
)

func FormatDuration(d time.Duration) string {
	db := DurationBreakdown{Input: d}

	const daysInMonth = 30 // Approximate

	x := 365 * 24 * time.Hour
	db.Years = d / x
	d -= db.Years * x

	x = daysInMonth * 24 * time.Hour
	db.Months = d / x
	d -= db.Months * x

	x = 24 * time.Hour
	db.Days = d / x
	d -= db.Days * x

	x = time.Hour
	db.Hours = d / x
	d -= db.Hours * x

	x = time.Minute
	db.Minutes = d / x
	d -= db.Minutes * x

	x = time.Second
	db.Seconds = d / x

	db.Output = ""

	if db.Years > 0 {
		db.Output += fmt.Sprintf("%dy", db.Years)
	}
	if db.Months > 0 {
		db.Output += fmt.Sprintf("%dM", db.Months)
	}
	if db.Days > 0 {
		db.Output += fmt.Sprintf("%dd", db.Days)
	}
	if db.Years == 0 {
		if db.Hours > 0 {
			db.Output += fmt.Sprintf("%dh", db.Hours)
		}
		if db.Days == 0 {
			if db.Minutes > 0 {
				db.Output += fmt.Sprintf("%dm", db.Minutes)
			}
			if db.Hours == 0 {
				if db.Seconds > 0 || db.Output == "" {
					db.Output += fmt.Sprintf("%ds", db.Seconds)
				}
			}
		}
	}

	slog.Debug(db.String())
	return db.Output
}
