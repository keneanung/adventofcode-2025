package day1

import (
	"fmt"
	"strconv"
	"strings"
)

// Step represents a movement step on a circular dial or wheel.
type Step struct {
	Direction       Direction
	Clicks          int
	EffectiveClicks int
	Position        int
	Wraps           int
}

const (
	startingValue  = 50
	numberOfValues = 100
)

// Direction represents the direction of movement on the dial.
type Direction byte

const (
	Left  Direction = 'L'
	Right Direction = 'R'
)

// GetSteps parses input lines and converts them into a sequence of steps representing movements
// on a circular dial or wheel. Each input line should contain a direction character ('L' for left
// or 'R' for right) followed by a number indicating the number of clicks to move.
//
// The function validates each input line to ensure:
// - The line has at least 2 characters
// - The first character is a valid direction ('L' or 'R')
// - The remaining characters form a valid non-negative integer
//
// For each valid input line, it calculates:
// - The effective clicks (clicks modulo the total number of values on the dial)
// - The new position after movement
// - The number of complete wraps around the dial
//
// Parameters:
//   - input: A slice of strings where each string represents a movement command
//
// Returns:
//   - []Step: A slice of Step structs containing movement details and calculated positions
//   - error: An error if any input line is invalid, with details about the line number and issue
//
// The function maintains state of the current position throughout the parsing process,
// starting from a predefined starting value and updating after each step.
func GetSteps(input []string) ([]Step, error) {
	currentPosition := startingValue
	result := make([]Step, 0, len(input))
	for index, line := range input {
		// normalize and validate input
		line = strings.TrimSpace(line)
		if len(line) < 2 {
			return nil, fmt.Errorf("line %d: not enough characters", index+1)
		}
		dir := Direction(line[0])
		if dir != Left && dir != Right {
			return nil, fmt.Errorf("line %d: invalid direction", index+1)
		}
		numberOfClicks, err := strconv.Atoi(line[1:])
		if err != nil || numberOfClicks < 0 {
			return nil, fmt.Errorf("line %d: invalid number of clicks", index+1)
		}

		effectiveClicks := numberOfClicks % numberOfValues
		delta := numberOfClicks
		if dir == Left {
			delta = -delta
		}
		end := currentPosition + delta
		newPosition := ((end % numberOfValues) + numberOfValues) % numberOfValues

		var wraps int
		if dir == Right {
			wraps = (currentPosition + numberOfClicks) / numberOfValues
		} else {
			wraps = (numberOfClicks + (numberOfValues - currentPosition)) / numberOfValues
			if currentPosition == 0 {
				wraps--
			}
			if wraps < 0 {
				wraps = 0
			}
		}

		result = append(result, Step{
			Direction:       dir,
			Clicks:          numberOfClicks,
			EffectiveClicks: effectiveClicks,
			Position:        newPosition,
			Wraps:           wraps,
		})
		currentPosition = newPosition
	}
	return result, nil
}
