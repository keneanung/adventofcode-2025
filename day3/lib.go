package day3

import (
	"fmt"
	"strconv"
)

// GetJoltage computes the sum of the largest subsequence numbers formed by
// selecting 'numberOfBatteries' digits from each line of input while preserving
// order. It validates inputs and returns the sum as int64.
func GetJoltage(input []string, numberOfBatteries int) (int64, error) {
	result := int64(0)
	for inputIndex, line := range input {
		if len(line) < numberOfBatteries {
			return 0, fmt.Errorf("line %d: line too short for %d batteries", inputIndex, numberOfBatteries)
		}
		for _, b := range []byte(line) {
			if b < '0' || b > '9' {
				return 0, fmt.Errorf("line %d: invalid character '%c' in line", inputIndex, b)
			}
		}
		batteryDigits := make([]byte, 0, numberOfBatteries)
		lastLargestDigitIndex := -1
		lineBytes := []byte(line)
		for missingDigits := numberOfBatteries - 1; missingDigits >= 0; missingDigits-- {
			// earliest index we can start searching for the next largest digit
			largestDigitIndex := lastLargestDigitIndex + 1
			largestDigit := lineBytes[largestDigitIndex]
			// latest index we can search to find the next largest digit
			searchWindowStart := largestDigitIndex + 1
			// leave at least i characters (the number of missing digits) at the end to reach the target length
			searchWindowEnd := len(lineBytes) - missingDigits
			for i:= searchWindowStart; i < searchWindowEnd; i++ {
				char := lineBytes[i]
				if char > largestDigit {
					largestDigit = char
					largestDigitIndex = i
				}
				if largestDigit == '9' {
					break
				}
			}
			batteryDigits = append(batteryDigits, largestDigit)
			lastLargestDigitIndex = largestDigitIndex
		}
		num, err := strconv.ParseInt(string(batteryDigits), 10, 64)
		if err != nil {
			return 0, fmt.Errorf("line %d: %w", inputIndex, err)
		}
		result += num
	}
	return result, nil
}

// GetJoltageStack computes the same result as GetJoltage, but uses a
// monotonic stack to build the maximum subsequence of length numberOfBatteries
// for each input line while preserving order. It validates inputs and returns
// the sum of those numbers as int64.
func GetJoltageStack(input []string, numberOfBatteries int) (int64, error) {
	result := int64(0)
	for inputIndex, line := range input {
		if len(line) < numberOfBatteries {
			return 0, fmt.Errorf("line %d: line too short for %d batteries", inputIndex, numberOfBatteries)
		}
		lineBytes := []byte(line)
		// Validate all characters are digits once up-front.
		for _, b := range lineBytes {
			if b < '0' || b > '9' {
				return 0, fmt.Errorf("line %d: invalid character '%c' in line", inputIndex, b)
			}
		}

		// Monotonic decreasing stack of digits (as bytes). We pop smaller digits
		// when a larger digit appears, as long as we can still reach length k.
		k := numberOfBatteries
		stack := make([]byte, 0, len(lineBytes))
		remaining := len(lineBytes)
		for _, d := range lineBytes {
			// While we have a smaller top and can still fill to k with remaining
			// digits (including current), pop.
			for len(stack) > 0 && stack[len(stack)-1] < d && (len(stack)-1+remaining) >= k {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, d)
			remaining--
		}

		// Take the first k digits from the stack to form the number.
		if len(stack) < k {
			// Should not happen due to earlier checks, but guard anyway.
			return 0, fmt.Errorf("line %d: insufficient digits after processing", inputIndex)
		}
		selected := stack[:k]
		// Convert selected digits to int64 without intermediate string allocations.
		var num int64
		for _, b := range selected {
			num = num*10 + int64(b-'0')
		}
		result += num
	}
	return result, nil
}
