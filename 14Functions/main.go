package main

import "fmt"

func foo() {
	fmt.Println("Hello from GO")
}

func add(val1 int, val2 int) int {
	return val1 + val2
}

// receive any amount of values
func proAdder (values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}

func main() {
	adder := add(5,10)
	fmt.Println(adder)

	proResult := proAdder(2,3,4,5,6,6)
	fmt.Println(proResult)
}
