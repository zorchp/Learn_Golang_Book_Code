package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex
	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("incrementing :%d\n", count)
	}
	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("decrementing :%d\n", count)
	}

	var arithemic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithemic.Add(1)
		go func() {
			defer arithemic.Done()
			increment()
		}()
	}
	for i := 0; i <= 5; i++ {
		arithemic.Add(1)
		go func() {
			defer arithemic.Done()
			decrement()
		}()
	}

	arithemic.Wait()
	fmt.Println("over")
}
