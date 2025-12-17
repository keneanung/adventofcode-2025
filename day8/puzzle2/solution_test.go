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
			input: []string{"162,817,812",
				"57,618,57",
				"906,360,560",
				"592,479,940",
				"352,342,300",
				"466,668,158",
				"542,29,236",
				"431,825,988",
				"739,650,466",
				"52,470,668",
				"216,146,977",
				"819,987,18",
				"117,168,530",
				"805,96,715",
				"346,949,466",
				"970,615,88",
				"941,993,340",
				"862,61,35",
				"984,92,344",
				"425,690,689",
			},
			expected: 25272,
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
