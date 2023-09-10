package tidehead

import (
	"fmt"
	"strings"
	"time"
)

type DurationBreakdown struct {
	Input   time.Duration
	Years   time.Duration
	Months  time.Duration
	Days    time.Duration
	Hours   time.Duration
	Minutes time.Duration
	Seconds time.Duration //lint:ignore ST1011 ignore this!
	Output  string
}

func (db DurationBreakdown) String() string {
	var sb strings.Builder

	sb.WriteString(db.Input.String())
	sb.WriteString(", ")
	sb.WriteString(fmt.Sprintf("y: %d", db.Years))
	sb.WriteString(", ")
	sb.WriteString(fmt.Sprintf("mo: %d", db.Months))
	sb.WriteString(", ")
	sb.WriteString(fmt.Sprintf("d: %d", db.Days))
	sb.WriteString(", ")
	sb.WriteString(fmt.Sprintf("h: %d", db.Hours))
	sb.WriteString(", ")
	sb.WriteString(fmt.Sprintf("m: %d", db.Minutes))
	sb.WriteString(", ")
	sb.WriteString(fmt.Sprintf("s: %d", db.Seconds))
	sb.WriteString(" gives ")
	sb.WriteString(db.Output)

	return sb.String()
}
