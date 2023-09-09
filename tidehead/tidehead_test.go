package main

import (
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		duration time.Duration
		expected string
	}{
		{3*365*24*time.Hour + 2*time.Hour + 15*time.Minute, "3y"},
		{3*24*time.Hour + 2*time.Hour + 15*time.Minute, "3d2h"},
		{2*24*time.Hour + 15*time.Minute + 55*time.Second, "2d"},
		{3*365*24*time.Hour + 2*time.Hour + 15*time.Minute, "3y"},
		{2*24*time.Hour + 15*time.Minute + 17*time.Second, "2d"},
		{3*24*time.Hour + 7*time.Hour + 59*time.Minute + 0*time.Second, "3d7h"},
		{3*24*time.Hour + 23*time.Minute + 0*time.Second, "3d"},
		{1*22*time.Hour + 27*time.Minute + 0*time.Second, "22h27m"},
		{2*365*24*time.Hour + 3*time.Hour, "2y"},
		{1*365*24*time.Hour + 6*time.Hour + 45*time.Minute + 30*time.Second, "1y"},
		{2*time.Hour + 1*time.Minute + 30*time.Second, "2h1m"},
		{33*24*time.Hour, "1M3d"},
		{33*24*time.Hour + 27*time.Minute + 10*time.Second, "1M3d"},
		{33*24*time.Hour + 10*time.Second, "1M3d"},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			result := FormatDuration(test.duration)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}