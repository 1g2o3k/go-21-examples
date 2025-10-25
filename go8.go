package main

import "fmt"

// add adds two numbers
func add(a, b int) int {
	return a + b
}

// subtract subtracts b from a
func subtract(a, b int) int {
	return a - b
}

// multiply multiplies two numbers
func multiply(a, b int) int {
	return a * b
}

// divide divides a by b
func divide(a, b int) int {
	return a / b
}

func main() {
	var a int
	fmt.Println("Enter number 1:")
	fmt.Scanf("%d", &a)
	var b int
	fmt.Println("Enter number 2:")
	fmt.Scanf("%d", &b)
	sum := add(a, b)
	difference := subtract(a, b)
	product := multiply(a, b)
	quotient := divide(a, b)
	fmt.Println("Sum:", a, "+", b, "=", sum)
	fmt.Println("Difference:", a, "-", b, "=", difference)
	fmt.Println("Product:", a, "*", b, "=", product)
	fmt.Println("Quotient:", a, "/", b, "=", quotient)
}
