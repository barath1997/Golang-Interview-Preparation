package main

import (
	"fmt"
	"strconv"
)

func main() {

	f := func(a int) string {
		b := strconv.Itoa(a)
		return b + b
	}

	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("array : %v\n", Map(arr, f))
}

func Map[T, U any](arr []T, fu func(T) U) []U {

	st := []U{}

	for _, val := range arr {
		st = append(st, fu(val))
	}

	return st

}
