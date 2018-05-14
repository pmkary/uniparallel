
# University Parallel

UniParallel is a simple atomic go library that provides you with some goodies in order to develop parallel software.

## What you get?

### Forall (from Multi Pascal)
In Multi-Pascal there is this grammar `forall` that is a parallel for. UniParallel provides the same functionality by a function called `ForAll`.

The function gets a range `start` to `end` for the iterating range, grouping of the threads a function (`func (index int) { ... }`) to be the body of the for.

```go
parallel.ForAll(start, end, grouping, func(index int) {
    fmt.Println( index )
})
```

### Spin Lock
UniParallel provides a Spin Lock implementation for Go. The lock implementation is fairly easy to use:

```go
var lock parallel.SpinLock
counter := 0
parallel.ForAll(1, 10, 4, func(index int) {
    lock.Lock()
    counter++
    lock.Unlock()
})
```

A higher level abstract is to use `RunSafe`. It locks the runtime and then runs a given function and unlocks as soon as the function ends. It's a very bad way to use the spin lock, very high footprint and so bad for the performance, but it's a pretty abstract:

```go
parallel.RunSafe(func() {
    someFunctionRunningInLockedRuntime()
})
```

### Terminal & I/O
As UniParallel is intended for use in a university course it has some easy functionality to handle data and terminal I/O to make it easier to write educational software

