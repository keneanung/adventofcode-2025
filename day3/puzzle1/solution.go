package puzzle1

import (
	"github.com/keneanung/adventofcode-2025/day3"
)

func Solve(input []string) (int64, error) {
	numberOfBatteries := 2
	return day3.GetJoltage(input, numberOfBatteries)
}
