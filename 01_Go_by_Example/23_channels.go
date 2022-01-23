package main

import "fmt"

//	Channels are pipes that connect concurrent goroutines.
//	You can send values into channels from one goroutine and receive those values into another goroutine

func main() {

	//	Create a new channel with make(chan type). Channels are typed by the values they convey
	messages := make(chan string)

	//	Send a value into a channel using the <- syntax
	//	Here we send 'ping' to the messages  channel we made above from a goroutine
	go func() { messages <- "ping" }()

	//	The <- channel syntax receives a value from the channel
	//	Here we'll receive the "ping" message we sent above and print it out
	msg := <-messages
	fmt.Println(msg)
}

//	When we run the program the "ping" message is successfully passed from one goroutine to another via our channel
//	By default, sends and receives block until both the sender and receiver are ready.
//	This property allows us to wait at the end of our program for the "ping" message without having to use any other synchronization