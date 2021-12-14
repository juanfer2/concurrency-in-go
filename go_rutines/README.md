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

De manera más abstracta las go rutines funcionan como corutines que son subprocesos que se ejecutan en y en ciertos puntos pueden ser suspendidas, reanudadas y canceladas. Lo que hace especial a las gorutines es su profunda intregración con el tiempo de ejecución que estas no definen sus puntos de suspensión o reentrada. Go observa su comportamiento en tiempo de ejecución y las suspende autómaticamente cuando las bloquean o las reanuda cuando las desbloquean, dicho de otra forma tiene interuptures para estos casos. Esto es una elegante asociación entre el tiempo de ejecución y la lógica de una gorutine.
