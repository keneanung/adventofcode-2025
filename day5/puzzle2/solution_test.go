package puzzle2

import (
	"testing"
)

func TestSolve(t *testing.T) {
	testInput := []struct {
		name     string
		input    []string
		expected int64
	}{
		{
			name: "example",
			input: []string{"3-5",
				"10-14",
				"16-20",
				"12-18",
				"",
				"1",
				"5",
				"8",
				"11",
				"17",
				"32"},
			expected: 14,
		},
	}

	for _, tt := range testInput {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Solve(tt.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}