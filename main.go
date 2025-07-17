package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Problem:
// - return 2 indices of the 2 numbers that add upt to the target.
// - there is only one solution.
// - cannot use the same element twice.

// First thoughts:
// - create logic for constraints.
// - use an auxiliar slice with all the numbers where their value is lesser than the target.
// - test all the possibilities using the smaller slice.

// Approach:
// - the filter approach was no good, since I didn't account for the fact that
// the numbers can be negative or zero.
// - brute force is an option but is not optimal.

func searchIndexes(nums []int, target int) []int {
	count := 0
	solution := make([]int, 2)
	for index, value := range nums {
		if value < int(math.Pow(-10, 9)) || value > int(math.Pow(10, 9)) {
			return []int{}
		}
		if value == target {
			solution[count] = index
			count++
		}
		if count >= 2 {
			return solution
		}
	}

	for index, value := range nums {
		for j := index + 1; j < len(nums); j++ {
			if value+nums[j] == target {
				solution[0] = index
				solution[1] = j
				break
			}
		}
	}

	return solution
}

func twoSum(nums []int, target int) []int {
	lenArr := len(nums)
	for _, value := range nums {
		if value < int(math.Pow(-10, 9)) || value > int(math.Pow(10, 9)) {
			return nil
		}
	}

	switch {
	case lenArr < 2 || lenArr > int(math.Pow(10, 4)):
		return []int{}

	case target < int(math.Pow(-10, 9)) || target > int(math.Pow(10, 9)):
		return []int{}

	default:
		return searchIndexes(nums, target)
	}
}

func printSolution(res []int) {
	strRes := make([]string, 2)
	for i, v := range res {
		strRes[i] = strconv.Itoa(v)
	}
	fmt.Println("[", strings.Join(strRes, ", "), "]")
}

func main() {
	fmt.Println("1. Two Sum")

	firstCase := []int{2, 7, 11, 15}
	res := twoSum(firstCase, 9)
	printSolution(res)

	secondCase := []int{3, 2, 4}
	res = twoSum(secondCase, 6)
	printSolution(res)

	thirdCase := []int{3, 3}
	res = twoSum(thirdCase, 6)
	printSolution(res)

	fourthCase := []int{0, 4, 3, 0}
	res = twoSum(fourthCase, 0)
	printSolution(res)

	fifthCase := []int{-3, 4, 3, 90}
	res = twoSum(fifthCase, 0)
	printSolution(res)

}
