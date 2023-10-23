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

func Selection(main [6]int) [6]int {
	temp := 0
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
			temp = main[x]
			main[x] = min
			main[index] = temp
		}
	}
	return main
}
