package main

import "fmt"

func greet2() {
	fmt.Println("Good afternoon, buddy")
}

func main() {
	fmt.Println("Function example")
	greet()
	greet2()
}

func greet() {
	fmt.Println("Good morning, everyone")
}
