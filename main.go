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
	pi := 3.14 //without specify type, compiler will do it for us

	fmt.Println(count, age, pi)

	// if contiditon
	if age >= 30 {
		fmt.Println("You are old")
	} else if age < 30 {
		fmt.Println("You are young")
	}

	// array / slice
	numbers := [5]int{4, 7, 1, 5, 2}                  // fixed size array
	strings := []string{"aa", "bb", "cc", "dd", "ee"} // a slice

	strings = append(strings, "ff") // you can modify later

	fmt.Println(numbers)
	fmt.Printf("%q\n", strings)

	// map
	phones := make(map[string]int) // map[<key type>]<value type>

	phones["dave"] = 2012348866
	phones["katy"] = 2014327499

	fmt.Println(phones)

	// loops
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	j := 0 // using like while loop
	for j < 5 {
		fmt.Println("j is ", j)
		j++
	}

	/*
		- another loop usage with a slice (or array)
		- change index to key for using with a map
	*/
	for index, value := range strings {
		fmt.Println("index", index, "value", value)
	}

}
