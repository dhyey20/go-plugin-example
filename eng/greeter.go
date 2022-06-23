package main

import (
	"fmt"

	"greeter/pkg/greeter"
)

type greeting struct {
	config string
}

func (g greeting) GreetFunction() string {
	fmt.Println("Hello Universe")
	return "Hello Universe"
}

func New() greeter.GreeterInterface {
	var Greeter greeter.GreeterInterface
	Greeter = greeting{
		config: "test",
	}
	return Greeter
}
