package main

import (
	"fmt"
	"strconv"
)

// Numbers holds two numbers as strings
type Numbers struct {
	Number1 string
	Number2 string
}

// multiply multiplies two integers
func multiply(a, b int) int {
	return a * b
}

func main() {
	var num1 string
	fmt.Println("Enter number 1:")
	fmt.Scanf("%s", &num1)
	var num2 string
	fmt.Println("Enter number 2:")
	fmt.Scanf("%s", &num2)
	numbers := Numbers{Number1: num1, Number2: num2}
	intNum1, err := strconv.Atoi(numbers.Number1)
	if err != nil {
		fmt.Println("Error converting number1:", err)
		return
	}
	intNum2, err := strconv.Atoi(numbers.Number2)
	if err != nil {
		fmt.Println("Error converting number2:", err)
		return
	}
	product := multiply(intNum1, intNum2)
	fmt.Println("Product:", product)
}
