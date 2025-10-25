package main

import (
	"fmt"
	"os"
)

func main() {
	var fileName string
	fmt.Println("File name:")
	fmt.Scanln(&fileName)
	var content string
	fmt.Println("File content:")
	fmt.Scanln(&content)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
