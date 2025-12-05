package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type FreshRange struct {
	Start int64
	End   int64
}

func ParseFreshIdRanges(input []string) ([]FreshRange, int, error) {
	freshIds := make([]FreshRange, 0)
	var index int
	for idx, line := range input {
		index = idx
		if len(line) == 0 {
			break
		}
		if !strings.Contains(line, "-") {
			return nil, 0, fmt.Errorf("invalid input at line %d: %s (expect dash)", index, line)
		}
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return nil, 0, fmt.Errorf("invalid input at line %d: %s (expect two parts)", index, line)
		}
		start, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return nil, 0, fmt.Errorf("invalid input at line %d: %s (cannot parse start)", index, line)
		}
		end, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return nil, 0, fmt.Errorf("invalid input at line %d: %s (cannot parse end)", index, line)
		}
		if start > end {
			return nil, 0, fmt.Errorf("invalid input at line %d: %s (start greater than end)", index, line)
		}
		freshIds = append(freshIds, FreshRange{Start: start, End: end})
	}
	slices.SortFunc(freshIds, func(i FreshRange, j FreshRange) int {
		return int(i.Start - j.Start)
	})
	return freshIds, index, nil
}
