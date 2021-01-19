package barrier

type Barrier struct {
	n    int
	c    chan int
	cond *Cond
}

func New(n int) *Barrier {
	barrier := &Barrier{n: n}
	barrier.cond = newCond()
	barrier.c = make(chan int)
	barrier.Init()
	return barrier
}

func (barrier *Barrier) Init() {

	go func() {
		for i := 0; i < barrier.n; i++ {
			barrier.c <- i
		}
	}()
}

func (barrier *Barrier) Wait() {
	if barrier.n < 2 {
		return
	}
	select {
	case num := <-barrier.c:
		if num < barrier.n-1 {
			barrier.cond.Wait()
		} else {
			barrier.cond.Broadcast()
			barrier.Init()
		}
	}
}
