package main

import "fmt"

// Student represents a student
type Student struct {
	Name       string
	Age        int
	Department string
}

func main() {
	student := Student{Name: "Gökalp", Age: 20, Department: "YBS"}
	fmt.Println(student)
}
