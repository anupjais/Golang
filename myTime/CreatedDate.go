package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Created my own date")

	createdDate := time.Date(2024, time.July, 27, 13, 12, 12, 98, time.UTC)

	fmt.Println("My create date = ", createdDate)
	fmt.Println("My create date = ", createdDate.Local().Format("01-02-2006 15:04:05 Monday 530GMT"))
}
