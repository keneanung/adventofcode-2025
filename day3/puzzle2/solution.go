package puzzle2

import (
	"github.com/keneanung/adventofcode-2025/day3"	
)

func Solve(input []string) (int, error) {
	numberOfBatteries := 12
	return day3.GetJoltage(input, numberOfBatteries)
}

