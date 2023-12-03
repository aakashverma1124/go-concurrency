package main

import "fmt"

/*
DeadlockDueToSender

The below code contains the channel, which has zero buffer capacity.
If we pass the value into the channel, and there is no corresponding goroutine ready to receive the value,
it’ll throw a deadlock error.

Unbuffered channels require both a sender and a receiver to be ready at the same time for communication to happen.
*/

func DeadlockDueToSenderOnly() {
	channel := make(chan int)
	channel <- 1
}

/*
DeadlockDueToReceiver

The <-channel line is blocking indefinitely because there's no other goroutine attempting to
send a value into the channel. The program is waiting for a value to be sent on the channel,
but there's no other part of the program providing that value.

Unbuffered channels require both a sender and a receiver to be ready at the same time for communication to happen.
*/

func DeadlockDueToReceiverOnly() {
	channel := make(chan int)
	<-channel
}

/*
DeadlockUsingForRange

Closing the channel is necessary to inform the receiver that no more values are to be received
and that the loop should be terminated.

The for..range uses a channel to iterate over.
If we don’t explicitly close the channel,
the range will keep waiting to receive the values from the channel.

But when we stop sending the values over the channel without closing it,
the for..range continues to wait for further values, and the goroutine or the function
that holds the for..range gets stuck, unable to progress.
In a scenario like this, neither any other function nor any goroutine can proceed further, leading to a deadlock.
*/

func DeadlockUsingForRange() {
	channel := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			channel <- i
		}
	}()

	for val := range channel {
		fmt.Printf("%d", val)
	}
}

/*
DeadlockUsingSelect

The select function randomly selects one of the case statements to execute from all the available statements
that have data to process. If no case statement is ready to process, the default statement will be executed.

This code will throw a deadlock error because no case statement is ready to execute,
and there’s no default statement.
The select function will wait forever, and since there are no other goroutines,
the code will throw a deadlock error.
*/

func DeadlockUsingSelect() {
	channel := make(chan int)
	select {
	case <-channel:
	}
}
