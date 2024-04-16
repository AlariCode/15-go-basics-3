package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Println(key, value)
	}
}
