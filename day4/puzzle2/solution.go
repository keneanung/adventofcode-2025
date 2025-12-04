package puzzle2

import (
	"github.com/keneanung/adventofcode-2025/day4"
)

func Solve(input []string) (int64, error) {
	printingRoom := day4.SetupRoom(input)

	movableRolls := 0
	rollsMoved := -1

	for rollsMoved != 0 {
		rollsMoved = day4.GetMovableRolls(printingRoom)
		movableRolls += rollsMoved
		for coords, roll := range printingRoom {
			if roll.KnownToBeMovable {
				delete(printingRoom, coords)
			}
		}
	}

	return int64(movableRolls), nil
}
