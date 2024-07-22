package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Deep dive into the Time inside Go lang")

	presentTime := time.Now()

	fmt.Println("Present Time : ", presentTime)
	fmt.Printf("Present Time type : %T \n", presentTime)
	fmt.Println("\nPresent Time 1 : ", presentTime.Format("01-02-2006"))
	fmt.Printf("Present Time type : %T \n", presentTime)
	fmt.Println("\nPresent Time 2 : ", presentTime.Format("01-Feb-2006"))
	fmt.Printf("Present Time type : %T \n", presentTime)
	fmt.Println("\nPresent Time 3 : ", presentTime.Format("01-02-2006 Monday"))
	fmt.Printf("Present Time type : %T \n", presentTime)
	fmt.Println("\nPresent Time 4 : ", presentTime.Format("01-02-2006 MON"))
	fmt.Printf("Present Time type : %T \n", presentTime)
	fmt.Println("\nPresent Time 5 : ", presentTime.Format("01/02/2006 Friday"))
	fmt.Printf("Present Time type : %T \n", presentTime)

}
