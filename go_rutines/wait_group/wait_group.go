package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}

	wg.Add(1)
	go sayHello()
	wg.Wait()

	// Don't Work
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)

	for _, salutationItem := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutationItem)
		}()
	}
	wg.Wait()

	// Work good
	fmt.Println("Good array")
	for _, salutationGood := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutationGretting string) {
			defer wg.Done()
			fmt.Println(salutationGretting)
		}(salutationGood)
	}
	wg.Wait()
}
