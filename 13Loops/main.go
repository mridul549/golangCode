package main

import "fmt"

func main() {
	days := []string {"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	fmt.Println(days)

	// Loop 1
	// for i := 0; i < 7; i++ {
	// 	fmt.Println(days[i])
	// }

	// Loop 2- better 
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// Loop 3
	for index, day := range days {
		fmt.Println("Index ", index, " day ", day)
	}

	/*
		to avoid index, same can be done for day
		for _, day := range days {
			fmt.Println("day ", day)
		}
	*/


}
