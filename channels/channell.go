package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

// sending
// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("Processing number:", num)
// 		time.Sleep(1 * time.Second) // Simulate some processing time
// 	}

// }

// receiving
// func sum(results chan int, num1 int, num2 int) {
// 	result := num1 + num2
// 	results <- result
// }

// go routine synchronizer
// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("Working...")
// }

// email is recieveonly and done is sendonly
// func emailSender(emails <-chan string, done chan<- bool) {
// 	defer func() { done <- true }()
// 	for email := range emails {
// 		fmt.Println("Sending email to:", email)
// 		time.Sleep(1 * time.Second) // Simulate sending email
// 	}
// }

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "HEHEHE"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Value := <-chan1:
			fmt.Println("Received from chan1:", chan1Value)
		case chan2Value := <-chan2:
			fmt.Println("Received from chan2:", chan2Value)
		default:
			fmt.Println("No data received from either channel")
		}
	}

	// Example of a buffered channel
	// email := make(chan string, 10) // Buffered channel with capacity 10
	// done := make(chan bool)

	// go emailSender(email, done)

	// for i := 0; i < 10; i++ {
	// 	email <- fmt.Sprintf("%d@example.com", i)
	// }

	// fmt.Println("Sending emails...")
	// close(email)
	// <-done

	// results := make(chan int)
	// go sum(results, 10, 2)
	// res := <-results
	// fmt.Println("Sum:", res)

	// numChan := make(chan int)
	// go processNum(numChan)

	// numChan <- 42
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// messageChan := make(chan string)
	// messageChan <- "ping" // This will block the main goroutine until another goroutine reads from the channel
	// // Start a goroutine to read from the channel
	// message := <-messageChan
	// fmt.Println(message)
}
