package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func Increment(c *Counter) {
	fmt.Println("===========start counter==============")
	c.mu.Lock()
	fmt.Println("Block :(")
	defer c.mu.Unlock()
	fmt.Println("UnlockBlock :)")
	c.value++
	fmt.Println(c.value)
	fmt.Println("end counter")
}

func main() {
	c := Counter{}

	var listItem = []int{1, 2, 3, 4, 5, 6}
	for index, _ := range listItem {
		fmt.Println("Start range")
		fmt.Println(index)
		Increment(&c)
		fmt.Println("End range")
	}
}
