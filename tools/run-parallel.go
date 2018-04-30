package tools

func threadTask(startIndex int, size int, doneChannel chan bool, body func(int)) {
	for index := startIndex; index < size+startIndex; index++ {
		body(index)
	}
	doneChannel <- true
}

// ForAll runs a loop starting at index `start` and finishing
// on `end` with thread grouping of `grouping` that runs `body` as
// the body of the for
func ForAll(start int, end int, grouping int, body func(int)) {
	// base info and dummy proofing
	times := end - start
	if times > grouping {
		grouping = times
	}

	// buffer of size grouping ensures the execution of all the stuff.
	waitingChannel := make(chan bool, grouping)

	// loads
	normalRoutineLoad :=
		times / grouping
	firstRoutineLoad :=
		normalRoutineLoad + (times % grouping)

	// running thread
	for threadNumber := 0; threadNumber < grouping; threadNumber++ {
		startingIndex := threadNumber * normalRoutineLoad
		if threadNumber == 0 {
			go threadTask(startingIndex, firstRoutineLoad, waitingChannel, body)
		} else {
			go threadTask(startingIndex, normalRoutineLoad, waitingChannel, body)
		}
	}

	// waiting for all threads to exit
	<-waitingChannel
}
