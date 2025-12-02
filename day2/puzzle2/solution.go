package puzzle2

import (
	"github.com/keneanung/adventofcode-2025/day2"
	"strconv"
)

func Solve(input []string) (int, error) {
	return day2.AddInvalidIds(input, isRepeatedSequence)
}

func isRepeatedSequence(i int) bool {
	if i <= 0 {
		return false
	}
	indexString := strconv.Itoa(i)
	indexLen := len(indexString)
	for l := indexLen / 2; l >= 1; l-- {
		if indexLen%l != 0 {
			continue
		}
		stringStart := indexString[:l]
		repeated := true
		for start := l; start+l <= indexLen; start += l {
			sequence := indexString[start : start+l]
			if sequence != stringStart {
				repeated = false
				break
			}
		}
		if repeated {
			return true
		}
	}
	return false
}
