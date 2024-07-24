package main

import "fmt"

func main() {

	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3")
	fmt.Println("Starts")
	myDefer()
	fmt.Println("Ends")
	defer fmt.Println("Defer 11")
	defer fmt.Println("Defer 12")
	defer fmt.Println("Defer 13")
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
