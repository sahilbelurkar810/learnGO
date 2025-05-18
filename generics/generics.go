package main

import (
	"fmt"
)

// func printSlice[T any](items []T) {

// 	for _, item := range items {
// 		fmt.Println(item)
// 	}

// }

func printSlice[T int | string](items []T) {

	for _, item := range items {
		fmt.Println(item)
	}

}

type Stack[T any] struct {
	elements []T
}

func main() {
	stack := Stack[int]{
		elements: []int{1, 2, 3},
	}
	Mystack := Stack[string]{
		elements: []string{"Alice", "Bob", "Charlie"},
	}
	fmt.Println(stack.elements)
	fmt.Println(Mystack.elements)
	names := Mystack.elements
	ages := stack.elements
	// names := []string{"Alice", "Bob", "Charlie"}
	// ages := []int{25, 30, 35}
	printSlice(names)
	printSlice(ages)
}
