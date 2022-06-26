package main

import (
	"fmt"
)

func helloWorld(name string) string {
	return "Hello, world! My name is " + name + "."
}

func main() {

	fmt.Println(helloWorld("Go Lang"))

}
