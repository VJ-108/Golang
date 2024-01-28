package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Making Different Request")
	// GET()
	// POST()
	// POST_FORM()
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
	ByteCount, _ := responseString.Write(content)

	// fmt.Println(string(content))
	fmt.Println("ByteCount: ", ByteCount)
	fmt.Println(responseString.String())
}

func POST() {
	const myURL = "http://localhost:8000/post"
	requestBody := strings.NewReader(`{
		"coursename" : "Hello",
		"price" : 0,
		"platform" : "Hello.com"
	}`)
	response, err := http.Post(myURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func POST_FORM() {
	const myURL = "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("firstname", "Hello")
	data.Add("lastname", "Hi")
	data.Add("email", "Hello@gmail.com")

	response, err := http.PostForm(myURL, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
