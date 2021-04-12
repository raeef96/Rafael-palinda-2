package main

import "fmt"

/*
// This program should go to 11, but sometimes it only prints 1 to 10.
func main() {
	ch := make(chan int)
	go Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	//PROBLEM: the for statement can be faster than the the channel reading,
	//which result in closing the channel before all values have been read.
	close(ch)
}
// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		fmt.Println(n)
	}
}
*/

//solution
func Main() {
	ch := make(chan int)
	go Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}

	//By having an loop like this makes sure that we dont close the channel before its empty.
	for len(ch) != 0 {}
	close(ch)
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		fmt.Println(n)
	}
}
