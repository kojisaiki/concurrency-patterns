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

func fanin(ch1, ch2 <-chan Message) <-chan Message {
	new_ch := make(chan Message)
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
	for i := 0; i < 100; i++ {
		msg1 := <-ch
		fmt.Println(msg1.str)
		msg2 := <-ch
		fmt.Println(msg2.str)

		<-msg1.block
		<-msg2.block
	}
}
