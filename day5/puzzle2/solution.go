package puzzle2

import (
	"github.com/keneanung/adventofcode-2025/day5"
)

func Solve(input []string) (int64, error) {
	freshIds, _, err := day5.ParseFreshIdRanges(input)
	if freshIds == nil {
		return 0, err
	}

	for {
		changed := false
		for i := len(freshIds) - 1; i > 0; i-- {
			if freshIds[i].Start <= freshIds[i-1].End {
				if freshIds[i].End > freshIds[i-1].End {
					freshIds[i-1].End = freshIds[i].End
				}
				freshIds = append(freshIds[:i], freshIds[i+1:]...)
				changed = true
			}
		}
		if !changed {
			break

		}
	}
	result := int64(0)
	for _, fr := range freshIds {
		result += (fr.End - fr.Start + 1)
	}
	return result, nil
}
