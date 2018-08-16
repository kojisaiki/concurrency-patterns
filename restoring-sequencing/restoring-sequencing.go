package main

import (
	"fmt"
	"time"
)

type Message struct {
	str   string
	block chan int
}

func generator(msg string) <-chan Message {
	ch := make(chan Message)
	blockingStep := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- Message{fmt.Sprintf("%s %d", msg, i), blockingStep}
			time.Sleep(time.Second)
			blockingStep <- 1
		}
	}()
	return ch
}

func fanin(ch1, ch2 <-chan string) <-chan string {
	new_ch := make(chan string)
	go func() {
		for {
			new_ch <- <-ch1
		}
	}()
	go func() {
		for {
			new_ch <- <-ch2
		}
	}()
	return new_ch
}

func main() {
	ch := fanin(generator("Hello"), generator("Bye"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
