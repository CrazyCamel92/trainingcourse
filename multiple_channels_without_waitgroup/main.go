package main

import (
	"fmt"
)
// using a channel to sync the action of 2 go routines instead of a waiting group
func main() {
	c:= make(chan int)
	channel_done:= make(chan bool)

	//push int values into the channel;
	go func(){
		for i:=0;i<10;i++  {
			c<-i
		}
		channel_done<- true;
	}()
	go func(){
		for i:=0;i<10;i++  {
			c<-i
		}
		channel_done<- true;
	}()
	// pull values of the synchronization channel
	go func() {
		<-channel_done
		<-channel_done
		close(c);
	}()
	// the range loop will print the value inside n after the synchronization channel is empty
	for n:= range c{
		fmt.Println(n)
	}
}
