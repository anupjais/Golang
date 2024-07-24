package main

import "fmt"

func main() {

	emp := usr{"Ramesh", 12, 123.4, true}

	fmt.Println("Name : ", emp.name)
	emp.getStatus()
}

type usr struct {
	name    string
	exp     int
	Salary  float64
	Retired bool
}

func (u usr) getStatus() {
	fmt.Println("Is retired : ", u.Retired)
}
