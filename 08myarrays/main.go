package main

import "fmt"

func main()  {

/*

-Arrays are not heavily used in Go unlike other languages

-Arrays are fixed runtime we should manually change the length of the array in compile time

*/ 

	
	fmt.Println("Welcome to arrays in Go")

	var veggylist = [3]string{"potato","tomato","mushroom"}

	fmt.Println("The veggu list is:",veggylist)
	fmt.Println("The veggu list is:",len(veggylist))
}

