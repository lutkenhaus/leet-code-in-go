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

func filterNums(nums []int, target int) (map[int]int, int) {
	mapPossibilities := make(map[int]int)

	for index, value := range nums {
		if value < int(math.Pow(-10, 9)) || value > int(math.Pow(10, 9)) {
			return nil, -1
		}
		if value < target {
			mapPossibilities[index] = value
		}
	}

	return mapPossibilities, len(mapPossibilities)
}

func twoSum(nums []int, target int) []int {
	lenArr := len(nums)
	if lenArr < 2 || lenArr > int(math.Pow(10, 4)) {
		return []int{}
	}
	if target < int(math.Pow(-10, 9)) || target > int(math.Pow(10, 9)) {
		return []int{}
	}

	solution := make([]int, 2)
	mapPossibilities, lenPossibilities := filterNums(nums, target)

	for index, value := range mapPossibilities {
		for j := index + 1; j < lenPossibilities; j++ {
			if value+mapPossibilities[j] == target {
				solution[0] = index
				solution[1] = j
			}
		}
	}

	return solution
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

}
