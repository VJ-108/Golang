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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonData := []byte(`{
		"coursename": "ReactJs",
        "website": "Hello.in",
        "tags": ["web-dev","js"]
	}`)

	var courses course
	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("Json is Valid")
		json.Unmarshal(jsonData, &courses)
		fmt.Printf("%#v\n", courses)
	} else {
		fmt.Println("Json is not Valid")
	}

	//some cases where you want to add data to key values

	var OnlineCourse map[string]interface{}
	json.Unmarshal(jsonData, &OnlineCourse)
	fmt.Printf("%#v\n", OnlineCourse)

	for k, v := range OnlineCourse {
		fmt.Printf("Key is %v, Value is %v and Type is %T\n", k, v, v)
	}
}
