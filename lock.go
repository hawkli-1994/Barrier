package barrier


type Lock struct {
	internal chan bool
}

func NewLock() *Lock {
	lock := Lock{internal:make(chan bool)}
	go func() {
		lock.internal <- true
	}()
	return &lock
}

func (lock *Lock) Lock() {
	<-lock.internal
}

func (lock *Lock) Unlock() {
	go func() {
		lock.internal <- true
	}()
}