package puzzle2

import ()

func Solve(input []string) (int64, error) {
	length := len(input[0])
	var beams []int
	for row := range input {
		line := input[row]
		newBeams := make([]int, length)
		for col := range line {
			if line[col] == '.' {
				if beams != nil {
					newBeams[col] += beams[col]
				}
				continue
			}
			if line[col] == 'S' {
				newBeams[col] = 1
				continue
			}
			if line[col] == '^' && beams[col] > 0 {
				newBeams[col-1] += beams[col]
				newBeams[col+1] += beams[col]
			}
		}
		beams = newBeams
	}
	result := 0
	for _, v := range beams {
		result += v
	}
	return int64(result), nil
}
