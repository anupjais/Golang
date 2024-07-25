package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://www.speedtest.net/result/16539641878"

// const myUrl = "http://127.0.0.1:5500/DomEaxample.html"

func main() {
	fmt.Println("SpeedTest Web Request")
	// req, err :=
	fmt.Println("My Url : ", myUrl)

	result, _ := url.Parse(myUrl)

	fmt.Println("Result : ", result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

}
