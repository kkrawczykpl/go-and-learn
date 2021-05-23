package main

import "fmt"

const helloWorldPrefix = "Hello world, "

func Hello(name string) string {

	if(name == "") {
		return "Hello, world"
	}
	return helloWorldPrefix + name
}

func main() {
	fmt.Println(Hello("Go!"))
}