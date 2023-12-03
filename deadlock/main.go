package main

func main() {
	// Run one function at any time because first function will not let run other functions due to deadlock.
	DeadlockDueToSenderOnly()
	DeadlockDueToReceiverOnly()
	DeadlockUsingForRange()
	DeadlockUsingSelect()
}
