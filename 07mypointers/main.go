package main

import "fmt"

func main()  {
	num := 100

	ptr := &num // ptr stores the reference || adress value of num which is stored
	fmt.Println("Value of pointer is ",&num) // prints the adress of num

	fmt.Println("The *Pointer stored adress value",*ptr)

	*ptr = *ptr+num
	fmt.Println("The new *Pointer value is",*ptr)
	fmt.Println("The new Num value is",num)
	// The Actuall num value is also changed becoz *ptr is referenced to num just like stack and heap in js
}