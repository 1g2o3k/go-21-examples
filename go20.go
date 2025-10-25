package main

import (
	"fmt"
	"os"
)

func main() {
	var option int
	fmt.Println("1. Create folder\n2. Create file\n3. Delete file\n4. Delete folder")
	fmt.Scanf("%d", &option)
	var name string
	fmt.Println("Name:")
	fmt.Scanf("%s", &name)
	switch option {
	case 1:
		err := os.Mkdir(name, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", err)
		} else {
			fmt.Println("Folder created")
		}

	case 2:
		file, err := os.Create(name)
		if err != nil {
			fmt.Println("Error creating file:", err)
		} else {
			file.WriteString(name + "\n")
			fmt.Println("File created")
			file.Close()
		}

	case 3:
		err := os.Remove(name)
		if err != nil {
			fmt.Println("Error deleting file:", err)
		} else {
			fmt.Println("File deleted")
		}

	case 4:
		err := os.RemoveAll(name)
		if err != nil {
			fmt.Println("Error deleting folder:", err)
		} else {
			fmt.Println("Folder deleted")
		}

	default:
		fmt.Println("Invalid option")
	}
}
