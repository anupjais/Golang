package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a number between 1 to 10 : ")

	input, _ := reader.ReadString('\n')

	fmt.Println("The rating is : ", input)
	fmt.Printf("The rating typt is = %T \n", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The rating will increased by 1 = ", nunumRating+1)
	}
}
