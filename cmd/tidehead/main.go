package main

import (
	"fmt"
	"time"

	"github.com/taylormonacelli/tidehead"
)

func main() {
	duration := 2*24*time.Hour + 15*time.Minute + 55*time.Second
	fmt.Println(tidehead.FormatDuration(duration))

	duration = 10*time.Hour + 15*time.Minute + 55*time.Second
	fmt.Println(tidehead.FormatDuration(duration))
}
