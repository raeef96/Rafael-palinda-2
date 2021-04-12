package main

import "fmt"

/*
// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string)
	ch <- "Hello world!"
	fmt.Println(<-ch) //PROBLEM: the thread starts waiting before the actual message have got sent above.
}
*/

//solution
func Main() {
	ch := make(chan string)

	go func() {
		//doing it in an own thread makes sure the message gets sent on the channel
		//even though the main thread is waiting.
		ch <- "Hello world!"
	}()

	fmt.Println(<-ch)
}
