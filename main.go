package main

import (
	"fmt"
	sorting "project1/arrays"
	baseMath "project1/math"
)

func Closure(testCase, num1, num2 int) func() {
	num := num1 * num2
	return func() { fmt.Printf("Closure test %d = %d\n", testCase, num) }
}

func main() {
	fmt.Printf("\nFunctions Playground\n\n")
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
	Closure(1, 222, 2)()
	funcVar := Closure(2, 222, 3)
	funcVar()
	fmt.Println("千万不要搜索! ", somethinSKetchy)

	fmt.Printf("\nArrays vs Maps 部分/Playground\n\n")
	arr := [6]int{8, 5, 6, 3, 1, 2} //array size needs to be studied more: declaring a function that can receive an array without prior knowledge of the array size
	fmt.Println("Unsorted array : ", arr)
	fmt.Println("Sorted array : ", sorting.Selection(arr))

	var mp map[string]int = map[string]int{
		"完了":   3,
		"神经病":  666,
		"啥意思?": 9,
		"我爱你":  520,
		"我姓石, 无论何时, 与你相识我都值": 1,
		"海公牛": 8,
	}
	fmt.Println("Map print : ", mp)
	search, okOrNotOk := mp["完了"]
	fmt.Println("Checking for map variable availability: ", search, okOrNotOk)
	search, okOrNotOk = mp["haha"]
	fmt.Println("If doesn't exist : ", search, okOrNotOk)
}
