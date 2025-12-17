package puzzle1

import "github.com/keneanung/adventofcode-2025/day7"

func Solve(input []string) (int64, error) {
	_, result, err := day7.IterateInput(input, func(newBeams *[]int, beams []int, col int) bool {
		if beams[col] > 0 {
			(*newBeams)[col] = 0
			(*newBeams)[col-1] = 1
			(*newBeams)[col+1] = 1
			return true
		}
		return false
	})
	return result, err
}
