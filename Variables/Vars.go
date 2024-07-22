package main

import "fmt"

func main() {
	// fmt.Println("Variabels")

	var name string
	fmt.Println(name)
	fmt.Printf("Variable type : %T \n", name)

	var num int8 = 127
	fmt.Println(num)
	fmt.Printf("Number1 : %T \n", num)

	var num2 int16 = 32767
	fmt.Println(num2)
	fmt.Printf("Number2 : %T \n", num2)
}
