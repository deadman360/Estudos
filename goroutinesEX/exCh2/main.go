package main

import "fmt"

func main() {
	var soma int
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num := make(chan int)
	for _, valor := range slice {
		go func() {
			num <- valor
		}()
		soma += <-num
	}

	fmt.Println(soma)
}
