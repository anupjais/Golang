package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter a data as Input example.")

	read := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a data : ")
	input, _ := read.ReadString('\n')

	fmt.Println("The data is : ", input)
	fmt.Printf("The data-type is : %T \n", input)
}
