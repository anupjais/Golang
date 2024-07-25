package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://geeksforgeek.com"

func main() {
	fmt.Println("Webrequest")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The response type is : %T\n", response)
	// fmt.Println("The response is ", response)

	defer response.Body.Close()
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
