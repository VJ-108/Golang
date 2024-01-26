package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("File Handling")
	content := "This needs to go in a file"

	file, err := os.Create("./Hello.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length: ", length)
	file.Close()
	readFile("./Hello.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data in File: ",string(dataByte))
}
