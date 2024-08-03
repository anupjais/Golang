package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Welcome to web verb")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Status ContentLength : ", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Status ContentLength : ", response.ContentLength)
	fmt.Println("Content = ", string(content))
}
