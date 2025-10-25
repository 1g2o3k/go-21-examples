package main

import "fmt"

func main() {
	var score1 int
	fmt.Println("Enter score 1:")
	fmt.Scanf("%d", &score1)
	var score2 int
	fmt.Println("Enter score 2:")
	fmt.Scanf("%d", &score2)
	var score3 int
	fmt.Println("Enter score 3:")
	fmt.Scanf("%d", &score3)
	average := (score1 + score2 + score3) / 3
	if average < 50 {
		fmt.Println("You failed the class")
	} else if average >= 50 {
		fmt.Println("You passed the class")
	} else {
		fmt.Println("Error")
	}
}
