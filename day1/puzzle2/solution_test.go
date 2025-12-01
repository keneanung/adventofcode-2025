package puzzle2

import (
	"testing"
)

func TestSolve(t *testing.T) {
	// read input from test_intput.txt
	testInput := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "example",
			input: []string{"L68",
				"L30",
				"R48",
				"L5",
				"R60",
				"L55",
				"L1",
				"L99",
				"R14",
				"L82"},
			expected: 6,
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