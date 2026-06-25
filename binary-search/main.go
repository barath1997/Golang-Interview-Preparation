package main

// binary search
func main() {

	arr := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}
	l := len(arr) - 1
	s := 0
	numberToBeFound := 6

	for s <= l {
		if arr[s] == numberToBeFound {
			println("1", s)
			return
		} else if arr[l] == numberToBeFound {
			println("2", l)
			return
		} else if numberToBeFound > arr[(s+l)/2] {
			s = ((s + l) / 2) + 1
		} else if numberToBeFound < arr[(s+l)/2] {
			l = ((s + l) / 2) - 1
		} else if numberToBeFound == arr[(s+l)/2] {
			println("3", (s+l)/2)
			return
		}
	}
	println(-1)
}
