package main

import (
	"fmt"
	"time"
)

func main() {
	chT := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		chT <- "Operaçã concluida"
	}()
	select {
	case msg := <-chT:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("Timeout")
	}
}
