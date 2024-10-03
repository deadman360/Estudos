package main

//Ping pong, faça duas go routines que "joguem" ping pong.

import (
	"fmt"
	"time"
)

func Pong(pingC, pongC chan string) {
	for {
		ping := <-pingC
		fmt.Println(ping)

		pongC <- "pong"
		time.Sleep(time.Second)
	}
}
func Ping(pingC, pongC chan string) {
	for {
		pingC <- "ping"

		pong := <-pongC
		fmt.Println(pong)
		time.Sleep(time.Second)
	}
}
func main() {
	pongC := make(chan string)
	pingC := make(chan string)

	go Pong(pingC, pongC)
	go Ping(pingC, pongC)

	select {}
}
