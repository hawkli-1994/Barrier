package csp

import (
	"time"
)

type Barrier struct {
	n        int
	channels []chan int
}

func NewBarrier(n int) *Barrier {
	channels := make([]chan int, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan int)
	}
	barrier := Barrier{
		n:        n,
		channels: channels,
	}
	go func() {
		barrier.channels[0] <- 1
	}()
	return &barrier
}

func (barrier *Barrier) Wait() {
	c := barrier.channels[0]
	n := <-c
	if n != barrier.n {
		c <- n + 1
		<-barrier.channels[n]
	}
	close(barrier.channels[n-1])
}

func main() {
	n := 3
	b := NewBarrier(n)
	for i := 1; i < n + 1; i++ {
		go func(s int) {
			time.Sleep(time.Duration(s) * time.Second)
			b.Wait()
			println(time.Now().Second())
		}(i)
	}
	time.Sleep(time.Duration(5) * time.Second)
}