# GO RUTINES

Una go rutine es una función que se ejecuta al mismo tiempo (no necesariamente en paralelo), junto con otro código

```go
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

```
