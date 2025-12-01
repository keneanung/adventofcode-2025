package puzzle1

import (
	"github.com/keneanung/adventofcode-2025/day1"
)

func Solve(input []string) (int, error) {
	steps, err := day1.GetSteps(input)
	if err != nil {
		return 0, err
	}
	result := 0
	for _, step := range steps {
		if step.ResultValue == 0 {
			result++
		}
	}
	return result, nil
}
