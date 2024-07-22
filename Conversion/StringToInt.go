package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Data conversion example")

	read := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a data : ")

	input, _ := read.ReadString('\n')
	fmt.Println("The data is : ", input)

	num, rr := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if rr != nil {
		fmt.Println(rr)
	} else {
		fmt.Println("The number is : ", num)
	}
}
