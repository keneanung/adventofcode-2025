package puzzle2

import (
	"github.com/keneanung/adventofcode-2025/day1"
)

func Solve(input []string) (int64, error) {
	steps, err := day1.GetSteps(input)
	if err != nil {
		return 0, err
	}
	result := int64(0)
	for _, step := range steps {
		result += int64(step.Wraps)
	}
	return result, nil
}
