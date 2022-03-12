package clockcombinations

import (
	"fmt"
	"strconv"
)

func DeterminePossibleClockDisplay(numbers []string) int {
	if len(numbers) != 4 {
		panic("Should only contain 4 numbers")
	}

	clockCombinationsCount := 0

	for currentIndex, number := range numbers {
		for innerIndex := currentIndex + 1; innerIndex < len(numbers); innerIndex++ {
			hourNumberString := number + numbers[innerIndex]
			hourNumber, err := strconv.Atoi(hourNumberString)
			if err != nil {
				panic(fmt.Errorf("given number %s is not a number", hourNumberString))
			}

			if hourNumber <= 23 {
				if currentIndex == 0 && innerIndex == 1 {
					minuteNumberString := numbers[2] + numbers[3]
					minuteNumber, err := strconv.Atoi(minuteNumberString)
					if err != nil {
						panic(fmt.Errorf("given number %s is not a number", minuteNumberString))
					}

					if minuteNumber <= 59 {
						clockCombinationsCount++
					}
				}
			}
		}
	}
	return clockCombinationsCount
}
