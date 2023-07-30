package main

import "fmt"

// if the first letter is capital that means the variable is set to public
const LoginToken string = "aufhhfdf"

func main()  {
	var username string = "mridul"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var usernam string = "mridul"
	fmt.Println(usernam)

	var egint int = 40
	fmt.Println(egint)
	fmt.Printf("Variable is of type: %T \n", egint)

	// max value is 255- refer docs for more info on datatypes
	var eguint8 uint8 = 255
	fmt.Println(eguint8)
	fmt.Printf("Variable is of type: %T \n", eguint8)
	
	var egfloat float32 = 5.292312433532
	fmt.Println(egfloat)
	fmt.Printf("Variable is of type: %T \n", egfloat)
	
	// default value for undeclared variable is 0
	// int is an alias
	var egint1 int 
	fmt.Println(egint1)
	fmt.Printf("Variable is of type: %T \n", egint1)

	// implicit type
	// no need to mention datatype of variable
	var website = "flavr.tech"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style
	// walurus operator
	// cant be used globally- use only in methods or locally
	noOfUser := 4
	fmt.Println(noOfUser)
	
	userCount := 4
	fmt.Println((userCount))
}