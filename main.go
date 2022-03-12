package main

import (
	"fmt"

	clockcombinations "github.com/keepitsimpol/clock/internal/service/clockCombinations"
)

func main() {
	fmt.Printf("Possible clock display combinations: %d", clockcombinations.DeterminePossibleClockDisplay([]string{"1", "2", "3", "4"}))
}
