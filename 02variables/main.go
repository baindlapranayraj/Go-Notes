package main

import "fmt"

var LogginToken = "gibberishValue"
/*
Public decleration,make sure the first letter of variable name is capital for Public variable in Golang
walrus operator does not work outside of main func
*/

func main()  {

	var age int = 10
	fmt.Println("Variable is",age)
	fmt.Printf("The variable type is %T \n",age)


	dataLength := 25
	fmt.Println("Variable is",dataLength)
	fmt.Printf("The variable type is %T \n",dataLength)

	// ++++++++++++++++++Default values++++++++++++
   var anotherVar int;
   fmt.Println(anotherVar);/* unlike javascript go does not store any garbage value in this 
                                it provides default value this fmt prints 0
						   	*/
	
	
	fmt.Println(LogginToken)// Printing the public value
	fmt.Printf("The variable type pf public is %T \n",LogginToken)

}