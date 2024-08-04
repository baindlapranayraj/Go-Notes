package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to structs in Go")

	v1 := User{
		"Pranay Raj",
		19,
		true,
	}

	fmt.Println("The user struct value:", v1)
	fmt.Printf("The user struct full value: %+v \n", v1)

	fmt.Printf("The user Name is: %v \n The user Age: %v \n", v1.Name,v1.Age)

	// Anonomus Structs
	v2 := struct {
		pokemon string
		power int
	}{
		pokemon: "Charizad",
		power: 10,
	}
	fmt.Printf("The name of pokemon is %v \n and its power level is %v", v2.pokemon,v2.power)
}

type User struct {
	Name string
	Age int
	isLogged bool
}