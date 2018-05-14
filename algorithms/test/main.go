package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/pmkary/parallel/tools"
)

func main() {
	A := []int{
		1, 18, 20, 40, 18, 50, 20, 4, 18, 100, 6, 400, 18, 7, 18,
		300, 18, 5, 7, 18, 400, 90, 100, 8, 18, 20, 90, 18, 200,
	}
	n := len(A)

	for i := 0; i < n; i++ {
		A[i] = rand.Intn(20)
	}

	counter := 0
	x := 18

	var lock tools.SpinLock
	tools.ForAll(0, n-1, 4, func(index int) {
		if A[index] == x {
			lock.Lock()
			counter++
			lock.Unlock()
		}
	})

	fmt.Println("Found: " + strconv.Itoa(counter))
}
