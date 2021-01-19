package barrier

type Cond struct {
	wait chan int
}

func newCond() *Cond {
	return &Cond{
		wait: make(chan int),
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
