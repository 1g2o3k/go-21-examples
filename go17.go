package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var count string
	fmt.Println("How many files to create?")
	fmt.Scanf("%s", &count)
	var fileName string
	fmt.Println("File name:")
	fmt.Scanf("%s", &fileName)
	countInt, err := strconv.Atoi(count)
	if err != nil {
		fmt.Println("Error converting count:", err)
		return
	}
	for i := 0; i < countInt; i++ {
		suffix := strconv.Itoa(i)
		fullName := fileName + suffix + ".txt"
		_, err := os.Create(fullName)
		if err != nil {
			fmt.Println("Error creating file:", err)
		}
	}
}
