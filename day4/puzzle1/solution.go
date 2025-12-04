package puzzle1

import (
	"github.com/keneanung/adventofcode-2025/day4"
)

func Solve(input []string) (int64, error) {
	printingRoom := day4.SetupRoom(input)

	movableRolls := day4.GetMovableRolls(printingRoom)

	return int64(movableRolls), nil
}
