package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile (fileName string) {
	databyte, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println(databyte)
	fmt.Println(string(databyte))
}

func main() {
	content := "This is the CEO of Bistroverse!"

	file, err := os.Create("./Bistroverse.txt")

	if err != nil {
		panic(err)
	}

	lenght, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is: ", lenght)
	defer  file.Close()

	readFile("./Bistroverse.txt")
}
