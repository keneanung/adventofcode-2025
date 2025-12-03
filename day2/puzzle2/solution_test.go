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
			name:     "example",
			input:    []string{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"},
			expected: 4174379265,
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
