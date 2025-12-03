package puzzle1

import (
	"strconv"

	"github.com/keneanung/adventofcode-2025/day2"
)

func Solve(input []string) (int64, error) {
	return day2.AddInvalidIds(input, hasTwoRepeatedHalves)
}

func hasTwoRepeatedHalves(i int) bool {
	indexString := strconv.Itoa(i)
	if len(indexString)%2 == 1 {
		return false
	}
	mid := len(indexString) / 2
	return indexString[:mid] == indexString[mid:]
}
