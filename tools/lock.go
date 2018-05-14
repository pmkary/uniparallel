package tools

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

func (lock *SpinLock) tryLock() bool {
	return atomic.CompareAndSwapUint32(&lock.status, 0, 1)
}
