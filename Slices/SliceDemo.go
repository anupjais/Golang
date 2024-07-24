package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices..")

	// var even = []int8{2, 4, 6, 8, 10}

	// fmt.Println("Elements : ", even)
	// fmt.Printf("Types of elements : %T\n", even)
	// fmt.Println("Size of elements : \n", Size(even))

	var fruits = []string{"Banana", "Graps", "Apple"}
	fmt.Println("Fruits are : ", fruits)
	fmt.Printf("Typesof fruits is : %T\n", fruits)

	fruits = append(fruits, "Mango")
	fmt.Println("Fruits : ", fruits)

	// fruits = append(fruits[6:])
	// fmt.Println("Fruits : ", fruits)  // slice bounds out of range [6:4]

	// fruits = append(fruits[1:3])
	// fmt.Println("Fruits : ", fruits) // [Banana Graps]

	// fruits = append(fruits[:2])
	// fmt.Println("Fruits : ", fruits)  // [Apple Banana]

	// fruits = append(fruits[:1])
	// fmt.Println("Fruits : ", fruits) // [Apple]

	// fruits = append(fruits[:1])
	// fmt.Println("Fruits : ", fruits) // []

	sort.Strings(fruits)
	fmt.Println("Sorted fruits : ", fruits)

}
