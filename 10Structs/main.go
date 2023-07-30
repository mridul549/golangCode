package main

import "fmt"

// capital U, all properties are capital letter first
type User struct {
	Name   string
	Email  string
	Status bool
}

func main() {
	// no inheritence in golang
	// no super, no parent
	mridul := User{"Mridul", "mridulv.it.21@nitj.ac.in", true}
	fmt.Println(mridul)
	fmt.Println(mridul.Name)
	fmt.Println(mridul.Email)
	fmt.Println(mridul.Status)
}
