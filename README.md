## Updated version

func main() { 


	// Create a buffered channel with a queue size of 10
	cnp := make(chan func(), 10)

    // introduce a wait group
	var wg sync.WaitGroup

    for i := 0; i < 4; i++ {
		// Call the function concurrently four times, additionally adding a wait group
		wg.Add(1)
		go func() {
			// reduce the value by one from the wait group at the end of the execution of goroutine
			defer wg.Done()
			// Range over the channel and execute the function
			for f := range cnp {
				f()
			}
		}()
	}

    // Pass a value to the channel
	cnp <- func() {
		fmt.Println("HERE1")
	}

    // Close the channel
	close(cnp)

    // Wait for all goroutines to finish
	wg.Wait()

    // Print 'Hello'
	fmt.Println("Hello")
}
