package puzzle2

import (
	"github.com/keneanung/adventofcode-2025/day8"
)

func Solve(input []string) (int64, error) {
	var lastEdge day8.Edge
	_, err := day8.AssembleCircuitNetwork(input, func(edges []day8.Edge, circuits *[]day8.Circuit, loopFunc func(day8.Edge)) {
		for _, edge := range edges {
			lastEdge = edge
			loopFunc(edge)
			if len(*circuits) == 1 {
				break
			}
		}
	})
	if err != nil {
		return 0, err
	}
	return int64(lastEdge.From.X * lastEdge.To.X), nil
}
