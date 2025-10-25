package main

import "fmt"

// add adds two integers
func add(a, b int) int {
	return a + b
}

func main() {
	var a int
	fmt.Println("Enter variable A:")
	fmt.Scanf("%d", &a)
	var b int
	fmt.Println("Enter variable B:")
	fmt.Scanf("%d", &b)
	sum := add(a, b)
	if sum <= 20 {
		fmt.Println("Sum is less than or equal to 20:", sum)
	} else {
		fmt.Println("Sum is greater than 20:", sum)
	}
}
