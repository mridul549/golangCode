package main

import "fmt"

func main() {
	fmt.Println("Hello")
	defer fmt.Println("World")
	defer fmt.Println("World1")
	defer fmt.Println("World2")
	// defer statements are executed just before the function is about to return 
	// they are executed in a LIFO fashion
	// the eg would print world2 world1 world
}
