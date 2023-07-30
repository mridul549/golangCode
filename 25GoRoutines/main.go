package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}

var wg sync.WaitGroup // usually are pointers
var mut sync.Mutex
var count = 0

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
}

func main() {
	// Example 1
	// go greeter("hello")
	// greeter("world")

	websitelist := []string {
		"https://lco.dev",
		"https://github.com",
		"https://google.com",
		"https://facebook.com",
		"https://flavr.tech",
		"https://microsoft.com",
		"https://cashfree.com",
		"https://paytm.com",
	}

	for _, website := range websitelist {
		go getStatusCode(website)
		wg.Add(1) // keep on increasing the count as to how many processes are not yet finished
	}

	wg.Wait()
	fmt.Println(signals)
	fmt.Println(count)
}
