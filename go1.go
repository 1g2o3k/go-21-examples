package main

import "fmt"

// add adds two integers
func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(5, 6))
}
