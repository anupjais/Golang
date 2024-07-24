package main

import "fmt"

func main() {
	// fmt.Println("")

	result := add(5, 4)

	fmt.Println("The result is : ", result)

	proAdd := proAdder(1, 2, 3, 4, 4, 6)
	fmt.Println("The pro addition : ", proAdd)

	spcl, msg := special(1, 2, 3, 4, 5)
	fmt.Println("The special : ", spcl)
	fmt.Println("The message : ", msg)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(vals ...int) int {
	result := 0
	for _, val := range vals {
		result += val
	}
	return result
}
func special(vals ...int) (int, string) {
	result := 0
	for _, val := range vals {
		result += val
	}
	return result, " Addition completed."
}
