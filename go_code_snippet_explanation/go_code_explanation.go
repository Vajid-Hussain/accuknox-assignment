package main

import "fmt"

func main() {
	// create a buffered channel with a queue size of 10
	cnp := make(chan func(), 10)

	for i := 0; i < 4; i++ {
		// call the function concurrently four times
		go func() {
			// range over the channel and execute the function
			for f := range cnp {
				f()
			}
		}()
	}

	// pass a value to the channel
	cnp <- func() {
		fmt.Println("HERE1")
	}

	// print 'Hello' then exit before executing 'HERE1' function
	fmt.Println("Hello")
}

// 1. Explaining how the highlighted constructs work:

// Ans) Creates a buffered channel of size 10. A for loop iterates four times to trigger go routines.
// So, 4 goroutines concurrently check if there is any value in the channel.
// A function is sent to the channel, but it doesn't get executed because the main goroutine exits before the function execution.

// 2. Giving use-cases of what these constructs could be used for:

// Ans) This is a fan-out pattern where a generator adds value to a channel,
// and parallel processes (goroutines) fetch the function and execute it,
// allowing tasks to be completed up to 4 times faster due to the four goroutines.

// 3. What is the significance of the for loop with 4 iterations?

// Ans) Each iteration triggers a goroutine. By the end of the loop, there are four goroutines working in parallel with the main goroutine,
// reading and executing functions from the channel concurrently.

// 4. What is the significance of make(chan func(), 10)?

// Ans) It creates a shared memory space with a size of 10 as a buffered channel.
// Under the hood, a queue is created and referenced by the cnp variable.
// The channel can hold 10 functions at a time, making it quicker by reducing the waiting time for space compared to an unbuffered channel.

// 5. Why is “HERE1” not getting printed?

// Ans) The reason is that the main goroutine exits before the 'HERE1' function starts.
// Go provides a solution in the form of a wait group. By introducing a wait group and closing the channel at the end of the function,
// we can ensure that 'HERE1' gets printed.
