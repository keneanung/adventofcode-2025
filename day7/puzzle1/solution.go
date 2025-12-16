package puzzle1

func Solve(input []string) (int64, error) {
	beams := make([]bool, len(input[0]))
	result := 0
	for row := range input {
		line := input[row]
		for col := 0; col < len(line); col++ {
			if line[col] == 'S' {
				beams[col] = true
				continue
			}
			if line[col] == '^' && beams[col] {
				result++
				beams[col] = false
				beams[col-1] = true
				beams[col+1] = true
			}
		}
	}
	return int64(result), nil
}
