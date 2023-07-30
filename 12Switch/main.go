package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// making a random dice number generator
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
		case 1:
			fmt.Println("Dice is 1")
		case 2:
			fmt.Println("Dice is 2")
		case 3:
			fmt.Println("Dice is 3")
		case 4:
			fmt.Println("Dice is 4")
		case 5:
			fmt.Println("Dice is 5")
		case 6:
			fmt.Println("Dice is 6")
	}



}
