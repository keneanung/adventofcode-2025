package puzzle1

import (
	"fmt"
	"strconv"

	"github.com/keneanung/adventofcode-2025/day5"
)

func Solve(input []string) (int64, error) {
	freshIds, index, err := day5.ParseFreshIdRanges(input)
	if freshIds == nil {
		return 0, err
	}
	spoiledCount := int64(0)
	for idx := index + 1; idx < len(input); idx++ {
		line := input[idx]
		availableId, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid input at line %d: %s (cannot parse available id)", idx, line)
		}
		for _, fr := range freshIds {
			if availableId >= fr.Start && availableId <= fr.End {
				spoiledCount++
				break
			}
		}
	}
	return spoiledCount, nil
}
