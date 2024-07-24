package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers examples")
	// reader := bufio.NewReader(os.Stdin)

	var num *int
	// fmt.Print("Entrer a number : ")
	// input, _ := reader.ReadString('\n')

	// fmt.Println("The data is : ", input)
	fmt.Println("The value of pointer : ", num)
	fmt.Printf("The value of pointer : %T\n", num)
	fmt.Println("The value of pointer & : ", &num)
	fmt.Printf("The value of pointer & : %T\n", &num)
	fmt.Printf("The value of pointer * : %T\n", *num)
	fmt.Println("The value of pointer * : ", *num)

}
