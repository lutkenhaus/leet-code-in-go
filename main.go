package main

import (
	"fmt"
	"math"
	"strings"
)

func fizzBuzz(n int) []string {
	if n < 1 || n > int(math.Pow(10, 4)) {
		return []string{}
	}

	result := make([]string, n)

	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result[i-1] = "FizzBuzz"
		case i%3 == 0:
			result[i-1] = "Fizz"
		case i%5 == 0:
			result[i-1] = "Buzz"
		default:
			result[i-1] = fmt.Sprintf("%d", i)
		}
	}

	return result
}

func main() {
	fmt.Println("[", strings.Join(fizzBuzz(3), ", "), "]")
	fmt.Println("[", strings.Join(fizzBuzz(5), ", "), "]")
	fmt.Println("[", strings.Join(fizzBuzz(15), ", "), "]")
}
