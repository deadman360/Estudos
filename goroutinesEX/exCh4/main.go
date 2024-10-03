package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)

	go produtor(ch)
	consumidor(ch)
}

func produtor(ch chan int) {
	for i := 0; i <= 100; i++ {
		ch <- i
		fmt.Println("Produzindo: ", i)
	}
	close(ch)
}
func consumidor(ch chan int) {
	for i := 0; i <= 100; i++ {
		fmt.Println("Consumindo: ", <-ch)
		time.Sleep(time.Millisecond * 500)
	}
}
