package main

import (
	"fmt"
	"sync"
)

func main() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("hello from %v!\n", id)
	}

	const num = 5
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
