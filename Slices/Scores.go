package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println("Hello there")

	Scores := make([]int, 4)
	fmt.Println("Scores : ", Scores)

	Scores[0] = 500
	Scores[1] = 520
	Scores[2] = 700
	Scores[3] = 100
	Scores = append(Scores, 200, 400, 350)
	fmt.Println("Scores are : ", Scores)

	fmt.Println("Is Sorted scores : ", sort.IntsAreSorted(Scores))
	sort.Ints(Scores)
	var indx int = 2
	fmt.Println("Sorted scores : ", Scores)
	// fmt.Println("Is Sorted scores : ", sort.IntsAreSorted(Scores))
	fmt.Println("Sorted upto index : ", append(Scores[:indx]))
	// Scores = append(Scores[:indx], Scores[indx+1:]...)
	// fmt.Println("Sorted upto index : ", Scores)
	fmt.Println("Scores are : ", append(Scores[:indx], Scores[indx+1:]...))
	// fmt.Println("Sorted scores : ", append(Scores[indx+1:]...))
}
