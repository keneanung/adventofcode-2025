package puzzle2

import (
	"github.com/hashicorp/go-set/v3"
	"github.com/keneanung/adventofcode-2025/day2"
	"strconv"
)

func Solve(input []string) (int, error) {
	return day2.AddInvalidIds(input, RepeatedSequence)
}

func RepeatedSequence(i int, result int) int {
	indexString := strconv.Itoa(i)
	indexStringRunes := []rune(indexString)
	for l := len(indexString) / 2; l >= 1; l-- {
		if len(indexString)%l != 0 {
			continue
		}
		chunked := make([]string, len(indexString)/l)
		for start := 0; start+l <= len(indexString); start += l {
			sequence := indexStringRunes[start : start+l]
			chunked[start/l] = string(sequence)
		}
		chunkSet := set.From(chunked)
		if chunkSet.Size() == 1 {
			result += i
			return result
		}
	}
	return result
}
