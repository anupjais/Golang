package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Take input from user")

	read := bufio.NewReader(os.Stdin)

	fmt.Print("Enter you name : ")

	// input, _ := read.ReadByte()
	// input, _ := read.ReadString()
	input, _ := read.ReadSlice(' ')
	fmt.Println("The input is ", input)
	fmt.Printf("The input type : %T\n", input)

}
