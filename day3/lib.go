package day3

import (
	"strconv"
)

func GetJoltage(input []string, numberOfBatteries int) (int, error) {
	result := 0
	for _, line := range input {
		batteryDigits := make([]byte, 0, numberOfBatteries)
		lastLargestDigitIndex := -1
		for i := numberOfBatteries - 1; i >= 0; i-- {
			largestDigit := line[lastLargestDigitIndex+1]
			largestDigitIndex := lastLargestDigitIndex + 1
			for index, char := range []byte(line[lastLargestDigitIndex+2 : len(line)-i]) {
				if char > largestDigit {
					largestDigit = char
					largestDigitIndex = index + lastLargestDigitIndex + 2
				}
				if largestDigit == '9' {
					break
				}
			}
			batteryDigits = append(batteryDigits, largestDigit)
			lastLargestDigitIndex = largestDigitIndex
		}
		num, err := strconv.Atoi(string(batteryDigits))
		if err != nil {
			return 0, err
		}
		result += num
	}
	return result, nil
}
