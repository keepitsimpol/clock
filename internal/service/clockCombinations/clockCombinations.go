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
		for itIndex := currentIndex; itIndex < len(numbers); itIndex++ {
			innerIndex := 0
			if itIndex+1 < len(numbers) {
				innerIndex++
			}

			hourNumberString := number + numbers[innerIndex]
			fmt.Printf("Current hour number: %s\n", hourNumberString)
			hourNumber, err := strconv.Atoi(hourNumberString)
			if err != nil {
				panic(fmt.Errorf("given number %s is not a number", hourNumberString))
			}

			if hourNumber <= 23 {
				fmt.Println("hour number is valid")
				if currentIndex == 0 && innerIndex == 1 {
					minuteNumberString := numbers[2] + numbers[3]
					fmt.Printf("Current minute number: %s\n", minuteNumberString)

					minuteNumber, err := strconv.Atoi(minuteNumberString)
					if err != nil {
						panic(fmt.Errorf("given number %s is not a number", minuteNumberString))
					}

					if minuteNumber <= 59 {
						fmt.Printf("Clock combination is: %s\n", hourNumberString+":"+minuteNumberString)
						clockCombinationsCount++
					} else {
						fmt.Println("Minute number is invalid")
					}

					minuteNumberString = numbers[3] + numbers[2]
					fmt.Printf("Current minute number: %s\n", minuteNumberString)

					minuteNumber, err = strconv.Atoi(minuteNumberString)
					if err != nil {
						panic(fmt.Errorf("given number %s is not a number", minuteNumberString))
					}

					if minuteNumber <= 59 {
						fmt.Printf("Clock combination is: %s\n", hourNumberString+":"+minuteNumberString)
						clockCombinationsCount++
					} else {
						fmt.Println("Minute number is invalid")
					}
				}

				if currentIndex == 1 && innerIndex == 2 {
					minuteNumberString := numbers[3] + numbers[0]
					fmt.Printf("Current minute number: %s\n", minuteNumberString)

					minuteNumber, err := strconv.Atoi(minuteNumberString)
					if err != nil {
						panic(fmt.Errorf("given number %s is not a number", minuteNumberString))
					}

					if minuteNumber <= 59 {
						fmt.Printf("Clock combination is: %s\n", hourNumberString+":"+minuteNumberString)
						clockCombinationsCount++
					} else {
						fmt.Println("Minute number is invalid")
					}

					minuteNumberString = numbers[0] + numbers[3]
					fmt.Printf("Current minute number: %s\n", minuteNumberString)

					minuteNumber, err = strconv.Atoi(minuteNumberString)
					if err != nil {
						panic(fmt.Errorf("given number %s is not a number", minuteNumberString))
					}

					if minuteNumber <= 59 {
						fmt.Printf("Clock combination is: %s\n", hourNumberString+":"+minuteNumberString)
						clockCombinationsCount++
					} else {
						fmt.Println("Minute number is invalid")
					}
				}

				if currentIndex == 2 && innerIndex == 3 {
					minuteNumberString := numbers[0] + numbers[1]
					fmt.Printf("Current minute number: %s\n", minuteNumberString)

					minuteNumber, err := strconv.Atoi(minuteNumberString)
					if err != nil {
						panic(fmt.Errorf("given number %s is not a number", minuteNumberString))
					}

					if minuteNumber <= 59 {
						fmt.Printf("Clock combination is: %s\n", hourNumberString+":"+minuteNumberString)
						clockCombinationsCount++
					} else {
						fmt.Println("Minute number is invalid")
					}

					minuteNumberString = numbers[1] + numbers[0]
					fmt.Printf("Current minute number: %s\n", minuteNumberString)

					minuteNumber, err = strconv.Atoi(minuteNumberString)
					if err != nil {
						panic(fmt.Errorf("given number %s is not a number", minuteNumberString))
					}

					if minuteNumber <= 59 {
						fmt.Printf("Clock combination is: %s\n", hourNumberString+":"+minuteNumberString)
						clockCombinationsCount++
					} else {
						fmt.Println("Minute number is invalid")
					}
				}

				if currentIndex == 3 && innerIndex == 0 {
					minuteNumberString := numbers[1] + numbers[2]
					fmt.Printf("Current minute number: %s\n", minuteNumberString)

					minuteNumber, err := strconv.Atoi(minuteNumberString)
					if err != nil {
						panic(fmt.Errorf("given number %s is not a number", minuteNumberString))
					}

					if minuteNumber <= 59 {
						fmt.Printf("Clock combination is: %s\n", hourNumberString+":"+minuteNumberString)
						clockCombinationsCount++
					} else {
						fmt.Println("Minute number is invalid")
					}

					minuteNumberString = numbers[2] + numbers[1]
					fmt.Printf("Current minute number: %s\n", minuteNumberString)

					minuteNumber, err = strconv.Atoi(minuteNumberString)
					if err != nil {
						panic(fmt.Errorf("given number %s is not a number", minuteNumberString))
					}

					if minuteNumber <= 59 {
						fmt.Printf("Clock combination is: %s\n", hourNumberString+":"+minuteNumberString)
						clockCombinationsCount++
					} else {
						fmt.Println("Minute number is invalid")
					}
				}
			} else {
				fmt.Println("Hour number is invalid")
			}
		}
	}
	return clockCombinationsCount
}
