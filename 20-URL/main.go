package main

import (
	"fmt"
	"net/url"
)

const myURL = "http://localhost:3000/learn?coursename=reactjs&paymentid=sdfsdfsf"

func main() {
	fmt.Println("URL Handling")
	fmt.Println(myURL)

	//Parsing the URL

	result, _ := url.Parse(myURL)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("Type of Query: %T\n", qParams)

	fmt.Println(qParams["coursename"])

	for _, v := range qParams {
		fmt.Println("Param: ", v)
	}

	partsOfUrl := &url.URL{
		Scheme:  "http",
		Host:    "localhost",
		Path:    "/course",
		RawPath: "user=Hello",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
