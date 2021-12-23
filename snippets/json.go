package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string
	Age    uint
	Gender string
}

func main() {
	var student = Student{
		Name:   "Krunal Vora",
		Age:    34,
		Gender: "male",
	}
	fmt.Printf("Student as a struct: %v", student)

	studentJson, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Student as JSON: %v", string(studentJson))
}
