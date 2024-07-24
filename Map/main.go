package main

import "fmt"

func main() {
	// fmt.Println("Maps here..")

	fmt.Println("\033[H\033[2J")
	language := make(map[string]string)

	language["JS"] = "JavaScript"
	language["PY"] = "Python"
	language["RB"] = "Ruby"
	language["KT"] = "Kotlin"
	language["GO"] = "Go Lang"

	fmt.Println("The datas are : ", language)
	// fmt.Printf("The datas type is : %T \n", language)

	// fmt.Println("\"Go\" stands for : ", language["GO"])
	// fmt.Println("\"KT\" stands for : ", language["KT"])

	// delete(language, "KT")
	// fmt.Println("The datas after deletion are : ", language)

	for key, value := range language {
		fmt.Printf("The key '%v', value is \"%v\"\n", key, value)
	}

}
