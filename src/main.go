package main

import (
	"fmt"
	"strings"
)

var (
	name string
	age  int
)

func main() {
	fmt.Println("Hello, World!")
	variable()
}

func variable() {
	fmt.Println("What is your name? ")
	name := "John Doe"
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s\n", name)

	var age, year = 18, 2023
	fmt.Printf("You are %d years old in %d?\n", age, year)
	var confirm string
	fmt.Scanln(&confirm)
	confirm = strings.ToLower(confirm)

	if confirm != "yes" {
		fmt.Println("What is your age? ")
		fmt.Scanln(&age)
	}
	fmt.Printf("You are %d years old in 2023!\n", age)
}
