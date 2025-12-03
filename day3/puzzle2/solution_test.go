package puzzle2

import (
	"testing"
)

func TestSolve(t *testing.T) {
	testInput := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "example",
			input: []string{"987654321111111",
				"811111111111119",
				"234234234234278",
				"818181911112111"},
			expected: 3121910778619,
		},
	}

	for _, tt := range testInput {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := Solve(tt.input)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
