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

	/*result, _ := url.Parse(myUrl)

	// fmt.Println("Result : ", result)
	fmt.Println("result.Scheme : ", result.Scheme)
	fmt.Println("result.Host : ", result.Host)
	fmt.Println("result.Port : ", result.Port())
	fmt.Println("result.Path : ", result.Path)
	fmt.Println("result.RawQuery : ", result.RawQuery)*/

	/*qpara := result.Query()

	fmt.Println("Query parameter : ", qpara)
	fmt.Printf("Query parameter : %T\n", qpara)*/

	// fmt.Println("Keys")
	// // fmt.Println(qpara["key"])
	// for _, val := range qpara {
	// 	fmt.Println("Parameters : ", val)
	// }
	// fmt.Println("Keys ends")

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "SpTest",
		Path:    "/result",
		RawPath: "user=user",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Another url : ", anotherUrl)

}
