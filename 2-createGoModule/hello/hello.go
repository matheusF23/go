package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Foo")
	fmt.Println(message)
}