package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i <= 20; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i <= 20; i++ {
			if i%2 != 0 {
				fmt.Println(i)
			}
		}
		wg.Done()
	}()

	wg.Wait()

}
