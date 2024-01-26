package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Making Different Request")
	GET()
}

func GET() {
	const myURL = "http://localhost:8000/get"
	response, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	ByteCount,_ := responseString.Write(content)

	// fmt.Println(string(content))
	fmt.Println("ByteCount: ",ByteCount)
	fmt.Println(responseString.String())

	
}
