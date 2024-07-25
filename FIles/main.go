package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("FIle system in Golang..")
	content := "Hey there I'm the content which is stored inside the file"

	file, err := os.Create("./myFile.txt")
	checkNil(err)

	length, err := io.WriteString(file, content)

	checkNil(err)
	fmt.Println("The length : ", length)

	defer file.Close()
	readFIle("./myFile.txt")

}

func readFIle(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNil(err)
	fmt.Println("The data in ASCII formate is : ", databyte)
	fmt.Println("The data is : ", string(databyte))
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
