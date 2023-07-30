package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://flavr.tech/products/getProductsByCategory?categoryName=All&outletid=64708814cb9778951511f8d3"

func main() {
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	for key, value := range qparams {
		fmt.Println(key, ": ", value)
	}

}
