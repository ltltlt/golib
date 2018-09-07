package main

import "fmt"

// Bisect search elem in array and return it's first position
func Bisect(array []int, elem int) int {
	// invariant: first elem in [l, u]
	l, u := 0, len(array)-1
	for l <= u {
		m := (l + u) / 2
		switch {
		case array[m] < elem:
			l = m + 1
		case array[m] == elem && (m == 0 || array[m-1] != elem):
			return m
		default:
			// array[m]>elem or array[m]==array[m-1]==elem
			u = m - 1
		}
	}
	return -1
}

func main() {
	pos := Bisect([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	fmt.Println(pos)
	pos = Bisect([]int{1, 2, 3, 3, 3, 4, 5}, 3)
	fmt.Println(pos)
	fmt.Println(Bisect([]int{3, 3}, 3))
}
