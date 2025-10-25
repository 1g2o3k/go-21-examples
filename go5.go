package main

import "fmt"

// Project represents a project
type Project struct {
	Name    string
	Subject string
	Time    int
}

func main() {
	project := Project{Name: "go file", Subject: "golang", Time: 12}
	fmt.Println(project.Name, "\n", project.Subject, "\n", project.Time)
}
