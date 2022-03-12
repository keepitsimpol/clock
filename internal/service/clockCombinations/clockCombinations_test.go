package clockcombinations

import "testing"

func TestDeterminePossibleClockDisplay(t *testing.T) {
	scenarios := []struct {
		testCaseName  string
		expectedCount int
		inputs        []string
	}{
		{
			testCaseName:  "1234",
			expectedCount: len([]string{"12:34", "12:43", "23:41", "23:14", "21:34", "21:43", "13:24", "13:42", "14:23", "14:32"}),
			inputs:        []string{"1", "2", "3", "4"},
		},
	}

	for _, tc := range scenarios {
		t.Run(tc.testCaseName, func(t *testing.T) {
			response := DeterminePossibleClockDisplay(tc.inputs)
			if response != tc.expectedCount {
				t.Errorf("Expected count: %d but was %d", tc.expectedCount, response)
			}
		})
	}
}
