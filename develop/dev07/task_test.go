package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	t.Run("First channel is closed", func(t *testing.T) {
		ch1 := make(chan interface{})
		go func() {
			time.Sleep(500 * time.Millisecond)
			close(ch1)
		}()

		ch2 := make(chan interface{})
		go func() {
			time.Sleep(1 * time.Second)
			close(ch2)
		}()

		done := or(ch1, ch2)

		doneTest := make(chan interface{})
		go func() {
			<-done
			close(doneTest)
		}()

		select {
		case <-doneTest:
		case <-time.After(2 * time.Second):
			t.Fatal("Expected at least one channel to be closed")
		}
	})

	// Add more test cases if necessary
}