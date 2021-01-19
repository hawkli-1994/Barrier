package barrier

type Cond struct {
	wait chan int
	L *Lock
}

func newCond(L *Lock) *Cond {
	return &Cond{
		wait: make(chan int),
		L: L,
	}
}

func (c *Cond) Wait() {
	select {
	case <-c.wait:
		return
	}
}

func (c *Cond) Broadcast() {
	close(c.wait)
	c.wait = make(chan int)
}
