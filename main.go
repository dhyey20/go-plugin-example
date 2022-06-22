package main

import (
	"fmt"
	"os"
	"plugin"
)

type GreeterInterface interface {
	GreetFunction() string
	// SecondGreetFunction() string
}

func main() {
	// determine module to load
	lang := "english"
	if len(os.Args) == 2 {
		lang = os.Args[1]
	}
	var mod string
	switch lang {
	case "english":
		mod = "./eng/eng.so"
	case "chinese":
		mod = "./chi/chi.so"
	case "swedish":
		mod = "./swe/swe.so"
	default:
		fmt.Println("don't speak that language")
		os.Exit(1)
	}

	// load module
	// 1. open the so file to load the symbols
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("after open")
	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	symNew, err := plug.Lookup("New")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	newFunction := symNew.(func())
	newFunction()
	fmt.Println("after new")

	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("after symgreeter")

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	var greeter GreeterInterface
	greeter, ok := symGreeter.(GreeterInterface)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	// 4. use the module
	returnString := (greeter).GreetFunction()
	fmt.Println("return string is ", returnString)
}
