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
			input: []string{".......S.......",
				"...............",
				".......^.......",
				"...............",
				"......^.^......",
				"...............",
				".....^.^.^.....",
				"...............",
				"....^.^...^....",
				"...............",
				"...^.^...^.^...",
				"...............",
				"..^...^.....^..",
				"...............",
				".^.^.^.^.^...^.",
				"...............",
			},
			expected: 40,
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
