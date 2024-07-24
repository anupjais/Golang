package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case statements")

	rand.Seed(time.Now().UnixNano())

	diceNum := rand.Intn(6) + 1
	fmt.Println("The dice number : ", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("You can open the game or move 1 spot")
	case 2:
		fmt.Println("You can move 2 spot")
		fallthrough
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot")

	}
	// fmt.Printf("The dice number : %T\n", diceNum)
	// fmt.Printf("The dice number : %v\n", diceNum)

}
