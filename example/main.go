package main

import (
	"fmt"
	"time"

	"github.com/go-softwarelab/common/pkg/seq"
	"github.com/go-softwarelab/common/pkg/seq2"
)

// Example of using seq and seq2 packages to make a simple clock with passed seconds time counter.
func main() {
	ticker := seq2.Tick(1 * time.Second)

	stringTime := seq2.MapValues(ticker, func(tick time.Time) string {
		return tick.Format("15:04:05")
	})

	messages := seq2.MapTo(stringTime, func(number int, tick string) string {
		return fmt.Sprintf("Current time %s\t\t\t You are waiting for %ds", tick, number)
	})

	seq.ForEach(messages, func(message string) {
		fmt.Print("\r", message)
	})
}
