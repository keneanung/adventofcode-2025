package day6

import (
	"fmt"
)

type Problem struct {
	Operand byte
	Values  []int
}

func GetResult(problems []Problem) (int64, error) {
	result := int64(0)
	for i, p := range problems {
		switch p.Operand {
		case '*':
			prod := int64(1)
			for _, v := range p.Values {
				prod *= int64(v)
			}
			result += prod
		case '+':
			sum := int64(0)
			for _, v := range p.Values {
				sum += int64(v)
			}
			result += sum
		default:
			return 0, fmt.Errorf("problem %d: unknown operand %q", i, p.Operand)
		}
	}
	return result, nil
}
