package main

import (
	"fmt"
	"sync"
	"time"
)

var contador = 0
var wg sync.WaitGroup
var mu sync.Mutex

func criarGoroutines(x int) {
	wg.Add(x)
	for i := 0; i < x; i++ {

		go func() {
			mu.Lock()
			z := contador
			z++
			contador = z
			time.Sleep(1000)
			fmt.Println(contador)
			wg.Done()
			mu.Unlock()
		}()
	}
}

func main() {

	criarGoroutines(10)
	wg.Wait()
}
