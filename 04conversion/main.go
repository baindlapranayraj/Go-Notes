package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	fmt.Println("Conversion in Go")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give your rating between 1 to 5")

	rating,_ := reader.ReadString('\n')
	fmt.Printf("The type of input is %T \n",rating);

	numRating,e := strconv.ParseFloat(strings.TrimSpace(rating),64);

	if e !=nil {
		fmt.Println(e)
	} else{
		fmt.Println("Added 1 to your rating",numRating+1)
	}
	
}

/*
   -- You might think this is a lot of work just for converting the string to number 
      But, we are able to handle or pridict the error with in line of code itself 
      We us alot like this for handling the API 
*/ 