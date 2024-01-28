package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //changes name to coursename in json formats
	price    int
	Platform string   `json:"website"` 
	Password string   `json:"-"` //This says that password will not be reflected in json
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs", 299, "Hello.in", "12345", []string{"web-dev", "js"}},
		{"MERN", 199, "Hello.in", "1234", []string{"web-dev", "js"}},
		{"Angular", 399, "Hello.in", "123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
