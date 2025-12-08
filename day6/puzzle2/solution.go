package puzzle2

import (
	"github.com/keneanung/adventofcode-2025/day6"
)

func Solve(input []string) (int64, error) {
	problems := make([]day6.Problem, 0)

	currentProblem := day6.Problem{}
	// read numbers top-down, so use the x axis as first loop
	for i := 0; i < len(input[0]); i++ {
		num := 0
		digitFound := false
		// read down the y axis, but skip the last line for now as this has the operand
		for j := 0; j < len(input)-1; j++ {
			char := input[j][i]
			if char != ' ' {
				digitFound = true
				num = num*10 + int(char-'0')
			}
		}
		// now set the operand from the last line
		if input[len(input)-1][i] != ' ' {
			currentProblem.Operand = input[len(input)-1][i]
		}

		if digitFound {
			// we found a number, add it to the current problem
			currentProblem.Values = append(currentProblem.Values, num)
		} else {
			// if the whole line was spaces only, we completed the last problem
			// we now need to store it and start a new one
			problems = append(problems, currentProblem)
			currentProblem = day6.Problem{}
		}
	}
	// append the last problem since we didn't have a whitespace only column at the end
	problems = append(problems, currentProblem)

	return day6.GetResult(problems)
}
