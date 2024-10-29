package main

import (
	"fmt"
	"strings"
)

func isValid(pattern string, start int) (bool, string) {
	nums := []int{start}

	var solve func(int) bool
	solve = func(pos int) bool {
		if pos == len(pattern) {
			return true
		}

		for next := 0; next <= 9; next++ {
			if isValidPair(nums[pos], next, pattern[pos]) {
				nums = append(nums, next)
				if solve(pos + 1) {
					return true
				}
				nums = nums[:len(nums)-1]
			}
		}
		return false
	}

	if !solve(0) {
		return false, ""
	}

	result := make([]string, len(nums))
	for i, n := range nums {
		result[i] = fmt.Sprintf("%d", n)
	}
	return true, strings.Join(result, "")
}

func isValidPair(prev, next int, symbol byte) bool {
	switch symbol {
	case 'L':
		return prev > next
	case 'R':
		return prev < next
	case '=':
		return prev == next
	}
	return false
}

func main() {
	encoded := []string{"LLRR=", "==RLL", "=LLRR", "RRL=R"}
	expected := []string{"210122", "000210", "221012", "012001"}

	for i, pattern := range encoded {
		for start := 0; start <= 9; start++ {
			if valid, result := isValid(pattern, start); valid {
				fmt.Printf("Pattern: %-6s Result: %-6s Expected: %s\n",
					pattern, result, expected[i])
				break
			}
		}
	}
}
