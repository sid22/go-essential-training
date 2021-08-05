package main

import (
	"fmt"
	"time"
)

/*
Channels in go are preferred way to communicate b/w go routines
channels are one direction pipes

If we try to receive from channel and it's empty we get blocked

For pushing, we have buffered and unbuffered channels

in unbuffered channels if we push we get blocked until someone picks it

In unbuffered channels we can push without getting blocked until the channel has empty space,
if channel is full and we push we will get blocked until someone picks one from the channel

*/

func main() {
	chanl := make(chan int)
	/*
		chanl := make(chan int)
		<-chanl // we will get blocked here
		fmt.Println("after channel")
	*/

	go func() {
		chanl <- 33
	}()

	val := <-chanl
	fmt.Printf("got %d\n", val)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending to chnl %d\n", i)
			chanl <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-chanl
		fmt.Printf("got from to chnl %d\n", val)
	}
}
