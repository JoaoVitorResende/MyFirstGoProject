package examplesfunctions

import (
	"fmt"
)

func SayHi() {
	fmt.Println("Hi")
}

func Add(a, b int) int {
	return a + b
}

func Swap(a, b int) (int, int) {
	return b, a
}

func Divide(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return res, rem
}

func AddHighOrder(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func AddVariant(nums ...int) int {
	var out int
	for _, n := range nums {
		out += n
	}
	return out
}
