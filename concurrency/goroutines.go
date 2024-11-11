package goroutines

import (
	"fmt"
	"time"
)

// Do simple add and push result into a channel
func Add(a, b int, result chan int) {
	sum := a + b
	result <- sum
}

func ChannelSynchronization(done chan bool) {
	fmt.Print("Before Sleep")
	time.Sleep(time.Second)

	done <- true
}
