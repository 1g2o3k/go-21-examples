package main

import (
	"fmt"
	"strconv"
)

func main() {
	var score1 string
	fmt.Println("Enter score 1:")
	fmt.Scanf("%s", &score1)
	var score2 string
	fmt.Println("Enter score 2:")
	fmt.Scanf("%s", &score2)
	var score3 string
	fmt.Println("Enter score 3:")
	fmt.Scanf("%s", &score3)
	intScore1, err := strconv.Atoi(score1)
	if err != nil {
		fmt.Println("Error converting score1:", err)
		return
	}
	intScore2, err := strconv.Atoi(score2)
	if err != nil {
		fmt.Println("Error converting score2:", err)
		return
	}
	intScore3, err := strconv.Atoi(score3)
	if err != nil {
		fmt.Println("Error converting score3:", err)
		return
	}
	average := (intScore1 + intScore2 + intScore3) / 3
	if average >= 50 {
		fmt.Println("Congratulations, you passed:", average)
	} else {
		fmt.Println("You failed:", average)
	}
}
