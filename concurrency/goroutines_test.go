package goroutines

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := make(chan int) //this is how channels made.

	go Add(3, 4, result) // this is how go routines are started

	sum := <-result // blocking pulling from channel

	expected := 7
	if sum != expected {
		t.Errorf("Expected %d, but %d", expected, sum)
	}
}

func TestChannelSynchronization(t *testing.T) {
	done := make(chan bool, 1)
	go ChannelSynchronization(done)

	<-done // this is blocked until something is pushed into.

}
