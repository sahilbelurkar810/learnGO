package main

func main() {
	// maps
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	// iterate over map
	for k, v := range m {
		println(k, v)
		if v == 2 {
			println("found 2")
		}
	}
}
