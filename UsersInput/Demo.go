package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for the Pizza : ")

	// comma, orr / error or
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us : ", input)
	fmt.Printf("Typeof this rating is %T\n", input)

	// Printf()
}
