package parallel

import (
	"runtime"
	"sync/atomic"
)

// SpinLock Structure
type SpinLock struct {
	status uint32
}

// Unlock unlocks the SpinLock
func (lock *SpinLock) Unlock() {
	atomic.StoreUint32(&lock.status, 0)
}

// Lock locks the SpinLock
func (lock *SpinLock) Lock() {
	for !lock.tryLock() {
		runtime.Gosched()
	}
}

// RunSafe gets a function (`body`) and runs it
// inside a locked and safe runtime time, it first
// locks the environment, exectues the code and
// then unlocks the spinlock
func (lock *SpinLock) RunSafe(body func()) {
	lock.Lock()
	body()
	lock.Unlock()
}

func (lock *SpinLock) tryLock() bool {
	return atomic.CompareAndSwapUint32(&lock.status, 0, 1)
}
