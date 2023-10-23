package main

import (
	"fmt"
	baseMath "project1/math"
)

func Return(testCase, num1, num2 int) func() {
	num := num1 * num2
	return func() { fmt.Printf("Closure test %d = %d\n", testCase, num) }
}

func main() {
	init := func(pi, radius float32) float32 {
		area := pi * radius * radius
		return area
	}(3.14, 5.00)

	somethinSKetchy := func(num1 int) int {
		out := num1 * 4
		out++
		return out
	}(44253)

	fmt.Println("Test circle area = ", init)
	fmt.Println("Function basic test = ", baseMath.Add(3, 4))
	Return(1, 222, 2)()
	funcVar := Return(2, 222, 3)
	funcVar()
	fmt.Println(somethinSKetchy)
}
