package main

import "fmt"

func sayHello() {
	fmt.Println("hello!!!")
}

func main() {
	go sayHello()

	fmt.Println("Run Anonimous Function")
	go func() {
		fmt.Println("Hi anonnimous function!!!")
	}()

	sayHi := func() {
		fmt.Println("Hi from sayHi")
	}
	go sayHi()
	// continue doing other things
}
