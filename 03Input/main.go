package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating!!")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input, " stars")


	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter rating: ")
	
	// // comma ok || err ok || input, err
	// // if not considering errors, we can take _
	// input, _ := reader.ReadString('\n')
	// fmt.Println("Thanks for rating us, ", input, " stars")
}