package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func AddInvalidIds(input []string, invalidIdFunc func(int, int) int) (int, error) {
	if len(input) != 1 {
		return 0, fmt.Errorf("expected single line input")
	}
	result := 0
	for idRange := range strings.SplitSeq(input[0], ",") {
		parts := strings.Split(idRange, "-")
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid range: %s", idRange)
		}
		startString := strings.TrimSpace(parts[0])
		start, err := strconv.Atoi(startString)
		if err != nil {
			return 0, fmt.Errorf("invalid start of range: %s", parts[0])
		}
		endString := strings.TrimSpace(parts[1])
		end, err := strconv.Atoi(endString)
		if err != nil {
			return 0, fmt.Errorf("invalid end of range: %s", parts[1])
		}
		for i := start; i <= end; i++ {
			result = invalidIdFunc(i, result)
		}
	}
	return result, nil
}
