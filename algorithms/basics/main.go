package main

import (
	"fmt"

	"github.com/pmkary/parallel/tools"
)

// main
func main() {
	// ForAll is a loop that has starting and ending
	// indexes, here that's from 2 to 10.
	//
	// Then it gets a grouping number which means how
	// many threads (go routines) you want to use for
	// this loop.
	//
	// At the end comes a function that takes an int
	// for the index and inside of it you specify what
	// you want as the body of your forloop to happen

	fmt.Println("This ForAll prints indexes from 2 to 10")
	tools.ForAll(2, 10, 4, func(index int) {
		fmt.Println(index)
	})
}
