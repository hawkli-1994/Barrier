package barrier

import (
	"fmt"
	"sync"
	"testing"
)

func TestBarrier_Wait(t *testing.T) {
	num := 5
	barrier := New(num)
	wg := sync.WaitGroup{}
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("A")
			barrier.Wait()			//barrier.Init()
			fmt.Println("B")
			barrier.Wait()
			fmt.Println("C")
		}()
	}
	wg.Wait()
}
