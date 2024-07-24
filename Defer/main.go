// defer behaves like STACK : Last In First Out
package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3")
	defer fmt.Println("Defer 4")
	defer fmt.Println("Defer 5")
	defer fmt.Println("Defer 6")

	fmt.Println("Hello there")
	defer fmt.Println("Defer 10")
	defer fmt.Println("Defer 9")

}
