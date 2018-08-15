package main

import (
	"fmt"
	"time"
)

func generator(msg string, waitFor time.Duration) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(waitFor * time.Millisecond)
		}
	}()
	return ch
}

func main() {
	ch1 := generator("Hello", 300)
	ch2 := generator("Bye", 1000)
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
	}
}
