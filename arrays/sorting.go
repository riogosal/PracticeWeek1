package sorting

import (
	baseMath "project1/math"
)

func SortedAscending(arr []int) bool {
	sort := true
	for i := 0; i < (len(arr) - 1); i++ {
		if arr[i] > arr[i+1] {
			sort = false
		}
	}
	return sort
}

func swap(arr []int, index1, index2 int) []int {
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp
	return arr
}

func Selection(main [6]int) [6]int {
	for x := 0; x < (len(main) - 1); x++ {
		min := 20
		index := x
		for y := x; y < len(main); y++ {
			current := main[y]
			min = baseMath.Min(current, min)
			if min == main[y] {
				index = y
			}
		}
		if index != x {
			main = [6]int(swap(main[:], x, index))
		}
	}
	return main
}
