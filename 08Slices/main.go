package main

import (
	"fmt"
	"sort"
)

func main()  {
	var fruitList = []string{"Apple", "Tomato"}

	// another easier way if we dont want to pass values on initialising
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("The fruit list is: ", fruitList)

	// some extra usecase of append function
	fruitList = append(fruitList[1:]) // [x:y] prints from x position to y, either of them can be blank like [:3] or [1:]
	fmt.Println(fruitList)

	// just another approach
	newList := make([]int, 4)
	for i := 0; i < 4; i++ {
		newList[i]=i+1
	}

	fmt.Println(newList)
	// now if we add another we may think that we'll recieve an error but no
	// if we use append method, we won't get any errors
	newList = append(newList, -1, -2, -4)
	fmt.Println(newList)

	// inbuilt sorting
	sort.Ints(newList)
	fmt.Println(newList)

	// remove a value from slices based on index 
	var index = 3
	newList = append(newList[:index], newList[index+1:]...)
	fmt.Println(newList)
}