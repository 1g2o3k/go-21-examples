package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	fmt.Println("Enter a:")
	fmt.Scanf("%s", &a)
	var b string
	fmt.Println("Enter b:")
	fmt.Scanf("%s", &b)
	intA, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Error converting a to int:", err)
		return
	}
	intB, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println("Error converting b to int:", err)
		return
	}
	sum := intA + intB
	fmt.Println("Sum:", sum)
}
