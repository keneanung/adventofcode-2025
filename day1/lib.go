package day1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Step struct {
	Direction   string
	Clicks      int
	ResultValue int
	ZeroPasses  int
}

const startingValue = 50
const numberOfValues = 100

type Direction byte

const (
	Left  Direction = 'L'
	Right Direction = 'R'
)

func GetSteps(input []string) ([]Step, error) {
	currentPosition := startingValue
	var result []Step
	for index, line := range input {
		// validate input
		if len(line) < 2 {
			return nil, fmt.Errorf("invalid input at line %d: Not enough characters", index+1)
		}
		r, size := utf8.DecodeRuneInString(line)
		if (r != rune(Left) && r != rune(Right)) {
			return nil, fmt.Errorf("invalid input at line %d: Invalid direction", index+1)
		}
		numberOfClicks, err := strconv.Atoi(line[size:])
		if err != nil || numberOfClicks < 0 {
			return nil, fmt.Errorf("invalid input at line %d: Invalid number of clicks", index+1)
		}

		var newPosition int
		zeroPasses := 0
		if Direction(r) == Left {
			newPosition = currentPosition - numberOfClicks
			if currentPosition == 0 {
				// if we start at 0 and go left, we did not pass zero but will count it below anyways
				// so make sure we have one less pass here
				zeroPasses--
			}
			for newPosition < 0 {
				newPosition = numberOfValues + newPosition
				zeroPasses++
			}
		} else if strings.HasPrefix(line, "R") {
			newPosition = currentPosition + numberOfClicks
			for newPosition >= numberOfValues {
				newPosition = newPosition - numberOfValues
				if newPosition != 0 {
					zeroPasses++
				}
			}
		}
		result = append(result, Step{
			Direction:   line[:1],
			Clicks:      numberOfClicks,
			ResultValue: newPosition,
			ZeroPasses:  zeroPasses,
		})
		currentPosition = newPosition
	}
	return result, nil
}
