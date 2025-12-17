package puzzle2

import "github.com/keneanung/adventofcode-2025/day7"

func Solve(input []string) (int64, error) {
	beams, _, err := day7.IterateInput(input, func(newBeams *[]int, beams []int, col int) bool {
		(*newBeams)[col-1] += beams[col]
		(*newBeams)[col+1] += beams[col]
		return false
	})
	if beams == nil {
		return 0, err
	}
	result := 0
	for _, v := range beams {
		result += v
	}
	return int64(result), nil
}
