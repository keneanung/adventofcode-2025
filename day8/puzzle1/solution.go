package puzzle1

import (
	"slices"

	"github.com/keneanung/adventofcode-2025/day8"
)

func Solve(input []string) (int64, error) {
	return RunIt(input, 1000)
}

func RunIt(input []string, numberOfConnections int) (int64, error) {
	circuits, err := day8.AssembleCircuitNetwork(input, func(edges []day8.Edge, circuits *[]day8.Circuit, loopFunc func(day8.Edge)) {
		for _, edge := range edges[:numberOfConnections] {
			loopFunc(edge)
		}
	})
	if err != nil {
		return 0, err
	}
	slices.SortFunc(circuits, func(c1 day8.Circuit, c2 day8.Circuit) int {
		return -1 * (c1.Size() - c2.Size())
	})
	result := 1
	for _, circuit := range circuits[:3] {
		result *= circuit.Size()
	}
	return int64(result), nil
}
