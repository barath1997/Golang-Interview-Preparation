package main

import (
	"cmp"
	"fmt"
)

type Comparator[T any] interface {
	Equal(a, b T) T
}

type IntComparator struct{}

func (it IntComparator) Equal(a, b int) bool {
	return a == b
}

type FloatComparator struct{}

func (ft FloatComparator) Equal(a, b float64) bool {
	return a == b
}

func main() {

	c := IntComparator{}
	fmt.Println("int comparator : ", c.Equal(5, 6))

	cf := FloatComparator{}
	fmt.Println("float comparator : ", cf.Equal(4.67, 4.67))

	///////////////////////////////////////

	fmt.Println(Sum(2, 3))
	fmt.Println(Sum(4.7, 3.8))
	fmt.Println(Sum("hello  ", "world"))
}

func Sum[T cmp.Ordered](a, b T) T {
	return a + b
}
