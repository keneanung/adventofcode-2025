package day4

type Coordinate struct {
	X int
	Y int
}

type Roll struct {
	KnownToBeMovable bool
}

type PrintingRoom map[Coordinate]Roll

func GetMovableRolls(printingRoom PrintingRoom) int {
	directions := []Coordinate{
		{X: 0, Y: -1},  // Up
		{X: 0, Y: 1},   // Down
		{X: -1, Y: 0},  // Left
		{X: 1, Y: 0},   // Right
		{X: -1, Y: -1}, // Up-Left
		{X: 1, Y: -1},  // Up-Right
		{X: -1, Y: 1},  // Down-Left
		{X: 1, Y: 1},   // Down-Right
	}
	movableRolls := 0
	for coords := range printingRoom {
		adjacentCount := 0
		for _, dir := range directions {
			adjacentCoord := Coordinate{X: coords.X + dir.X, Y: coords.Y + dir.Y}
			if _, exists := printingRoom[adjacentCoord]; exists {
				adjacentCount++
			}
		}

		if adjacentCount < 4 {
			movableRolls++
			printingRoom[coords] = Roll{KnownToBeMovable: true}
		}
	}
	return movableRolls
}

func SetupRoom(input []string) PrintingRoom {
	printingRoom := make(PrintingRoom)

	for y, line := range input {
		for x, char := range line {
			if char == '@' {
				printingRoom[Coordinate{X: x, Y: y}] = Roll{KnownToBeMovable: false}
			}
		}
	}
	return printingRoom
}
