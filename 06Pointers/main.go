package main

import "fmt"

func main() {
	fmt.Println("Hi")
	var ptr *int

	fmt.Println("Value of pointer is: ", ptr)

	var score int = 40

	var scorePtr = &score
	fmt.Println("Value of ptr2 is: ", *scorePtr)
	fmt.Printf("Type of ptr2 is: %T \n", scorePtr)

	*scorePtr *= 2
	fmt.Println("Value of ptr2 is: ", *scorePtr)
	fmt.Println("Value of score is: ", score)
	

}