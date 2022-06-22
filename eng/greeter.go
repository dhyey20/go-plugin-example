package main

import "fmt"

type greeting struct {
	config string
}

var Greeter greeting

func (g greeting) GreetFunction() string {
	fmt.Println("Hello Universe")
	return "Hello Universe"
}

func New() {
	Greeter = greeting{
		config: "test",
	}
}
