package main

import (
	"fmt"
	log "log/slog"
	"time"
)

func FormatDuration(d1 time.Duration) string {
	const daysInMonth = 30 // Approximate value

	years := d1 / (365 * 24 * time.Hour)
	d1 -= years * 365 * 24 * time.Hour
	months := d1 / (daysInMonth * 24 * time.Hour)
	d1 -= months * daysInMonth * 24 * time.Hour
	days := d1 / (24 * time.Hour)
	d1 -= days * 24 * time.Hour
	hours := d1 / time.Hour
	d1 -= hours * time.Hour
	minutes := d1 / time.Minute
	d1 -= minutes * time.Minute
	seconds := d1 / time.Second

	x := fmt.Sprintf("y: %d, mo: %d, d: %d, h: %d, m: %d, s: %d\n", years, months, days, hours, minutes, seconds)
	log.Debug(x)

	result := ""

	if years > 0 {
		result += fmt.Sprintf("%dy", years)
	}
	if months > 0 {
		result += fmt.Sprintf("%dM", months)
	}
	if days > 0 {
		result += fmt.Sprintf("%dd", days)
	}
	if years <= 0 {
		if hours > 0 {
			result += fmt.Sprintf("%dh", hours)
		}
		if days <= 0 {
			if minutes > 0 || result == "" {
				result += fmt.Sprintf("%dm", minutes)
			}
			if hours <= 0 {
				if seconds > 0 || result == "" {
					result += fmt.Sprintf("%ds", seconds)
				}
			}
		}
	}

	return result
}
