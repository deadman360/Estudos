package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)
	for i := 0; i <= 100; i++ {
		go func() {
			fmt.Println("Eu sou  GO FUNC DE ID: ", i)
			ch <- i
		}()
		output := <-ch
		fmt.Println(output)
	}

}
