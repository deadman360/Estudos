package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	dados := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int, len(dados))
	for _, v := range dados {
		wg.Add(1)
		go func(v int) {
			fmt.Println("adicionando:", v*v)
			ch <- v * v
			wg.Done()
		}(v)
	}
	var dadosQuadrados []int
	wg.Add(1)
	go func() {
		for i := 0; i < len(dados); i++ {
			dadosQuadrados = append(dadosQuadrados, <-ch)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(dadosQuadrados)
	fmt.Println(sum(dadosQuadrados))
}
func sum(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}
