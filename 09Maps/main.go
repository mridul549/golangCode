package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	// for deleting
	delete(languages, "JS")
	fmt.Println(languages)

	// looping on maps
	for key, value := range languages {
		fmt.Println("For key ", key, " the value is ", value)
		// fmt.Printf("For key %v, the value is %v", key, value)
	}
}
