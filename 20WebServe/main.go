package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func sendFormData () {
	// doing an example POST request, don't execute this

	urlToSend := "http://localhost:3005/postData"

	data := url.Values{}
	data.Add("firstName", "mridul")
	data.Add("lastName", "verma")
	data.Add("email", "mridulv.it.21@nitj.ac.in")

	response, err := http.PostForm(urlToSend, data)

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
	defer response.Body.Close()
}

func makePostReq() {
	// doing an example POST request, don't execute this

	url := "http://localhost:3005/postData"

	requestBody := strings.NewReader(`
		{
			"email": "mridulverma478@gmail.com"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	// any data that comes from the web is always in the form of bytedata, 
	// make sure to convert it to readable format before using it further
	fmt.Println(string(content))
}

func makeGetReq() {
	url := "https://flavr.tech/products/getProductsByCategory?categoryName=All&outletid=64708814cb9778951511f8d3"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func main() {
	makeGetReq()
}
