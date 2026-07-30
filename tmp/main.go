package main

import "fmt"

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Printf("arr : %v\n", replaceElements(arr))
}
func replaceElements(arr []int) []int {
	l := len(arr)
	if l < 2 {
		return nil
	}
	leader := arr[l-1]
	for i := l - 1; i >= 0; i-- {
		if arr[i] > leader {
			leader = arr[i]
		}
		arr[i] = leader
	}
	arr[l-1] = -1
	return arr
}
