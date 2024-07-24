package main

import "fmt"

func main() {
	fmt.Println("Loop example")

	days := []string{"SUN", "MON", "TUE", "THR", "RI", "SAT", "SAT"}

	fmt.Println("The daye are : ", days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Print(days[d], " ")
	// }
	// fmt.Println()

	// for i := range days {
	// 	fmt.Print(days[i], " ")
	// }

	// for i := 1; i < 5; i++ {
	// 	fmt.Print(i, " ")
	// }
	// fmt.Println()

	// for indx, day := range days {
	// 	fmt.Printf("%v day of days is \"%v\"\n", indx, day)
	// }
	// fmt.Println("")

	for _, day := range days {
		fmt.Printf("The day is \"%v\"\n", day)
	}
	fmt.Println("")

	/* 	//  Not working in Go lang
	for (var i int := 0; i < len(days); i++ ) {
		fmt.Print(ddays[i]," ")
	}
	fmt.Println("")
	*/

	num := 1

	for num < 20 {
		if num%2 != 0 {
			num++
			continue
		}
		fmt.Print(num, " ")
		num++
	}
	fmt.Println("")

}
