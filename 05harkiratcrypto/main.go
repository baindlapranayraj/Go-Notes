package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Crypto")

	read := bufio.NewReader(os.Stdin)
	input,_ := read.ReadString('\n')

	fmt.Println(input)
}