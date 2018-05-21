package parallel

import (
	"sync"
)

// Barrier implements a barrier to wait for
// all commands
type Barrier struct {
	c      int
	n      int
	m      sync.Mutex
	before chan int
	after  chan int
}

// NewBarrier creates a new Barrier of size n
func NewBarrier(n int) *Barrier {
	return &Barrier{
		n:      n,
		before: make(chan int, n),
		after:  make(chan int, n),
	}
}

// Before keeps go rutines waiting for the
// task
func (b *Barrier) Before() {
	b.m.Lock()
	b.c++
	if b.c == b.n {
		for i := 0; i < b.n; i++ {
			b.before <- 1
		}
	}
	b.m.Unlock()
	<-b.before
}

// After gets used after the barrier to make
// all the rutines wait
func (b *Barrier) After() {
	b.m.Lock()
	b.c--
	if b.c == 0 {
		for i := 0; i < b.n; i++ {
			b.after <- 1
		}
	}
	b.m.Unlock()
	<-b.after
}
