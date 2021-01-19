package barrier



type Barrier struct {
	n    int
	cur  int
	cond *Cond
}

func New(n int) *Barrier {
	barrier := &Barrier{n:n}
	barrier.Init()
	return barrier
}

func (barrier *Barrier) Init() {
	lock := NewLock()
	barrier.cond = newCond(lock)
	barrier.cur = barrier.n
}

func (barrier *Barrier) Wait() {
	barrier.cond.L.Lock()
	defer barrier.cond.L.Unlock()
	if barrier.cur--; barrier.cur > 0 {
		barrier.cond.Wait()
	} else {
		barrier.cond.Broadcast()
		barrier.cur = barrier.n
	}
}
