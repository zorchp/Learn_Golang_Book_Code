package main

import (
	"fmt"
	"sync"
)

// type Locker interface {
// 	Lock()
// 	Unlock()
// }

func main() {
	var mu sync.Mutex
	count := 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100_000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
