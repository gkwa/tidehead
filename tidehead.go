package tidehead

import (
	"fmt"
	"log/slog"
	"time"
)

func FormatDuration(d time.Duration) string {
	const daysInMonth = 30 // Approximate

	years := d / (365 * 24 * time.Hour)
	d -= years * 365 * 24 * time.Hour
	months := d / (daysInMonth * 24 * time.Hour)
	d -= months * daysInMonth * 24 * time.Hour
	days := d / (24 * time.Hour)
	d -= days * 24 * time.Hour
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute
	d -= minutes * time.Minute
	seconds := d / time.Second

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

	out := ""
	out += fmt.Sprintf("duration: %v", d)
	out += ", "
	out += fmt.Sprintf("y: %d", years)
	out += ", "
	out += fmt.Sprintf("mo: %d", months)
	out += ", "
	out += fmt.Sprintf("d: %d", days)
	out += ", "
	out += fmt.Sprintf("h: %d", hours)
	out += ", "
	out += fmt.Sprintf("m: %d", minutes)
	out += ", "
	out += fmt.Sprintf("m: %d", seconds)
	out += " gives "
	out += result

	slog.Debug(out)

	return result
}
