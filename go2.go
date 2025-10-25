package main

import "fmt"

// multiply multiplies the number by 100
func multiply(a int) int {
	return a * 100
}

func main() {
	for i := 0; i <= 1000; i++ {
		fmt.Println(multiply(101))
	}
}
