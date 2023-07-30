package main

import "fmt"

func main()  {
	var fruitList [4]string 
	
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("The fruit list is: ", fruitList)
	fmt.Println("The size of the array is: ", len(fruitList))
	
}