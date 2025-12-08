package puzzle1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/keneanung/adventofcode-2025/day6"
)

func Solve(input []string) (int64, error) {
	numberOfProblems := len(splitline(input[0]))
	problems := make([]day6.Problem, numberOfProblems)
	for lineNo, line := range input {
		split := splitline(line)
		var function func(*day6.Problem, string) error
		if lineNo == len(input)-1 {
			function = setOperand
		} else {
			function = appendValue
		}
		for i, value := range split {
			p := problems[i]
			err := function(&p, value)
			if err != nil {
				return 0, fmt.Errorf("line %d field %d: %w", lineNo, i, err)
			}
			problems[i] = p
		}
	}
	return day6.GetResult(problems)
}

func setOperand(p *day6.Problem, value string) error {
	p.Operand = value[0]
	return nil
}

func appendValue(p *day6.Problem, value string) error {

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	p.Values = append(p.Values, intValue)

	return nil
}

func splitline(line string) []string {
	return strings.FieldsFunc(line, func(r rune) bool {
		return r == ' '
	})
}
