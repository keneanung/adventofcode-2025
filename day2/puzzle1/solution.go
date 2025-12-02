package puzzle1

import (
	"strconv"

	"github.com/keneanung/adventofcode-2025/day2"
)

func Solve(input []string) (int, error) {
	return day2.AddInvalidIds(input, TwoSameSequences)
}

func TwoSameSequences(i int, result int) int {
	indexString := strconv.Itoa(i)
	if len(indexString)%2 == 1 {
		return result
	}
	mid := len(indexString) / 2
	if indexString[:mid] == indexString[mid:] {
		result += i
	}
	return result
}
