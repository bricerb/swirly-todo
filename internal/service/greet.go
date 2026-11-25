package service

import (
	"fmt"
	"todos/internal/io"
)

func Greet() User {
	user := User{}
	fmt.Print("Enter your name: ")

	input, err := io.TakeUserInput()

	if err != nil {
		panic(err)
	}

	fmt.Println("INPUT: ", input)
	user.Name = input

	fmt.Printf("Hello, %s! Fuck you beeeeetch! \n", user.Name)
	return user
}

