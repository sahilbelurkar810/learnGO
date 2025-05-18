package main

import "fmt"

func main() {
	// Iterate over array using range
	nums := []int{2, 3, 4}
	for i, num := range nums {
		fmt.Printf("index: %d, value: %d\n", i, num)
	}

	// Iterate over map using range
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Iterate over string using range
	for i, c := range "go" {
		fmt.Printf("index: %d, char: %c\n", i, c)
	}
}
