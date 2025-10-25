package main

import "fmt"

// Car represents a car
type Car struct {
	Model string
	Year  int
}

func main() {
	car := Car{Model: "renault", Year: 10}

	for i := 0; i <= 120; i++ {
		fmt.Println(car)
	}
}
