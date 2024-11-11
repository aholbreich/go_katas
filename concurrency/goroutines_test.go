package goroutines

import (
	"testing"
	"time"
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

// Showecase of the very nice Select feature
func TestSelectWithTimeout(t *testing.T) {

	ch1 := make(chan int)
	ch2 := make(chan int)

	// Start two goroutines with fixed delays
	go func() {
		time.Sleep(300 * time.Millisecond)
		ch1 <- 10
	}()
	go func() {
		time.Sleep(700 * time.Millisecond)
		ch2 <- 20
	}()

	select { //accepts what comes firts.useful for tasks like multiplexing, timeouts

	case result := <-ch1:
		expected := 10
		if result != expected {
			t.Errorf("Expected %d from ch1, but got %d", expected, result)
		}
	case result := <-ch2:
		t.Errorf("Did not expect to receive from ch2, but got %d", result)
	case <-time.After(500 * time.Millisecond):
		t.Error("Timeout occurred, but ch1 should have sent a value first")
	}
}
