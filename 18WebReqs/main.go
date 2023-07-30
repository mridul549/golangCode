package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func errorCheck (err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	res, err := http.Get(url)

	errorCheck(err)
	fmt.Printf("Type of req is: %T", res)

	// always close the connection manually, your responsibility
	defer res.Body.Close()

	databytes, err := ioutil.ReadAll(res.Body)

	errorCheck(err)
	content := string(databytes)
	fmt.Println("Content is: ", content)

}
