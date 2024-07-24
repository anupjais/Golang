package main

import "fmt"

func main() {
	// fmt.Println("Structs here")

	// emp := User{"Ramesh", "ramesh@ceo.com", 3, true}
	emp := User{"Ramesh", 12, 3, true}

	fmt.Println("The user is : ", emp)
	fmt.Printf("The user type is : %+v\n", emp)

	fmt.Printf("The name %v is %vyr old in experience with $%vk salary\n", emp.name, emp.age, emp.salary)
}

type User struct {
	name   string
	salary float64
	age    int
	status bool
}
