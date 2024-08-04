package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(`Welcome to time`)

	currentData := time.Now();

	fmt.Println(currentData.Format("01-02-2006 Monday"))
}