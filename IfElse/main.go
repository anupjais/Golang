package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Else if")

	// num := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter a number to check even or not : ")
	var num int = 4
	// input, _ := num.ReadString('\n')

	// newNum, err := strconv.ParseInt(strings.TrimSpace(input), 64)

	/*if err != nil {
		fmt.Println("No input data")
	} else {
	}
	}*/
	if num%2 == 0 {
		fmt.Println(num, " is and even")
	} else {
		fmt.Println(num, " is and even")
	}
}
