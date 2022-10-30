package practice4

import (
	"fmt"
)

func Factorial(n *int) int {
	sum := 1
	for i := 2; i <= *n; i++ {
		sum *= i
	}
	fmt.Println(n)
	return sum
}