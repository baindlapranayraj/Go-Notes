package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	fmt.Println("Welcome to input using bufio")

	reader := bufio.NewReader(os.Stdin) // buffio is used to read the user input
	fmt.Println("Enter Your age")

	// comma ok || comma error syntax similar to try-catch in /javascript

	input,_ := reader.ReadString('\n') //Reads the input
	fmt.Println("Your Age is",input) 
	fmt.Printf("The type of given input is %T \n",input)

}