package day7

func IterateInput(input []string, splitterCallback func(*[]int, []int, int) bool) ([]int, int64, error) {
	length := len(input[0])
	var beams []int
	counter := 0
	for row := range input {
		line := input[row]
		newBeams := make([]int, length)
		for col := range line {
			switch line[col] {
			case '.':
				if beams != nil {
					newBeams[col] += beams[col]
				}
			case 'S':
				newBeams[col] = 1
			case '^':
				resultedInSplit := splitterCallback(&newBeams, beams, col)
				if resultedInSplit {
					counter++
				}
			}
		}
		beams = newBeams
	}
	return beams, int64(counter), nil
}
