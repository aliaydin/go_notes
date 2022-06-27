package main

import (
	"fmt"
)

// a function
func helloWorld(name string) string {
	return "Hello, world! My name is " + name + "."
}

// we use main function to initiate the program
func main() {

	//call a function as a parameter in calling another function
	fmt.Println(helloWorld("Go Lang"))

	//define variables
	var count int
	var age int = 25
	pi := 3.14

	fmt.Println(count, age, pi)

	// contiditon
	if age >= 30 {
		fmt.Println("You are old")
	} else if age < 30 {
		fmt.Println("You are young")
	}

}
