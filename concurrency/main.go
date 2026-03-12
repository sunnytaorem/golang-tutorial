package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true // send a signal to the channel that the task is done
	close(doneChan)  // close the channel to signal that no more values will be sent. This only makes sense if we know which operation takes the longest. Because once the channel is close, it will stop sending data and goroutines will stop as well
}

// In this example, all functions doesn't have to run one after another. We can run them concurrently, without waiting for each other to finish.
// By running them concurrently, we can improve the overall performance and responsiveness of our program.
// We can do this by using the go keyword before the function call. This will create a new goroutine that will run the function concurrently with the main goroutine.
// In the example below, the main goroutine will continue to execute the next lines of code without waiting for the goroutines to finish.
// It means the main goroutine will invoke all the functions concurrently and it will stop, without waiting for the goroutines to finish. The output can be different each time.
// The solution to this is by using a channel. It serves as a communication mechanism between goroutines, allowing them to synchronize and coordinate their execution.
// In this example, we create a channel called doneChan of type bool. Each goroutine will send a signal (true) to the channel when it finishes its task.
// The main goroutine will wait for the signal from the channel using <-doneChan, which will block until it receives a value from the channel, ensuring that the main goroutine only continues after the slowGreet function has completed its execution.

func main() {
	// We can also create a channel to wait for completion of multiple goroutines
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// dones[0] = make(chan bool)
	// go greet("Nice to meet you!", dones[0])
	go greet("Nice to meet you!", done)
	// dones[1] = make(chan bool)
	// go greet("How are you?", dones[1])
	go greet("How are you?", done)

	// dones[2] = make(chan bool)
	// go slowGreet("How ... are ... you ...?", dones[2])
	go slowGreet("How ... are ... you ...?", done)
	// dones[3] = make(chan bool)
	// go greet("I hope you're liking the course!", dones[3])
	go greet("I hope you're liking the course!", done)

	// Wait for all goroutines to finish
	// for _, done := range dones {
	// 	<-done // wait for the signal from the channel that the goroutine is done.
	// }
	// go will only stop after data comes out of the channel
	for range done {
		// fmt.Println(doneChan)
	}

}
