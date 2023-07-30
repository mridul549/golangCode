package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // giving aliases to original names which will be returned in the response
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // '-' means don't include in the response
	Tags     []string `json:"tags,omitempty"` // doesn't show nils
}

func EncodeJson() {
	lcoCourse := []course{
		{"reactjs", 399, "lco.in", "abc123", []string{"web, mobile"}},
		{"nodejs", 199, "lco.in", "bcd123", []string{"web, mobile"}},
		{"nextjs", 399, "lco.in", "cde123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "reactjs",
			"Price": 399,
			"website": "lco.in",
			"tags": [ "web, mobile" ]
		}
	`)

	jsonValid := json.Valid(jsonDataFromWeb)
	var lcoCourse course
	if jsonValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON is not valid")
	}

	// some cases where you just want to add data to key value

	var myData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("%v: %v Type: %T \n", k, v, v)
	}

}

func main() {
	fmt.Println("Welcome to json video")
	DecodeJson()
}
