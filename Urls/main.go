package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://chat.qspiders.com/student-dashboard"

func main() {
	fmt.Println("Urls here..")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println("Result ", result)
	fmt.Println("Scheme ", result.Scheme)
	fmt.Println("port() ", result.Port())
	fmt.Println("path ", result.Path)
	fmt.Println("Host ", result.Host)
	fmt.Println("RawQuery ", result.RawQuery)
	fmt.Println("RawFragment ", result.RawFragment)
	fmt.Println("RequestURI() ", result.RequestURI())
	fmt.Println("RawPath ", result.RawPath)
}
