package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var option int
	fmt.Println("1. Create files\n2. Remove files")
	fmt.Scanf("%d", &option)
	switch option {
	case 1:
		var fileName string
		fmt.Println("File name:")
		fmt.Scanf("%s", &fileName)
		var amount int
		fmt.Println("Amount of files:")
		fmt.Scanf("%d", &amount)
		for i := 0; i < amount; i++ {
			suffix := strconv.Itoa(i)
			_, err := os.Create(fileName + suffix + ".txt")
			if err != nil {
				fmt.Println("Error creating file:", err)
			}
		}
	case 2:
		var nameToRemove string
		fmt.Println("File name to remove:")
		fmt.Scanf("%s", &nameToRemove)
		var amount int
		fmt.Println("Amount of files:")
		fmt.Scanf("%d", &amount)
		for i := 0; i < amount; i++ {
			suffix := strconv.Itoa(i)
			err := os.Remove(nameToRemove + suffix + ".txt")
			if err != nil {
				fmt.Println("Error removing file:", err)
			}
		}
	default:
		fmt.Println("Invalid option")
	}
}
