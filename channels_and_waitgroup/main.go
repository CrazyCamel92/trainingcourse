package main

import (
	"fmt"
	"sync"
)
// using range to handle the waiting for the go routine
func main() {
	c:= make(chan int)
	//creating and setting the wait group to wait for 2 dones
	var wg sync.WaitGroup;
	wg.Add(2);

	//push int values into the channel;
	go func(){
		for i:=0;i<10;i++  {
			c<-i
		}
	wg.Done();
	}()
	go func(){
		for i:=0;i<10;i++  {
			c<-i
		}
	wg.Done();
	}()
	// after the wait group is done the data will pass through the channel
	go func() {
		wg.Wait()
		close(c);
	}()
	// the range loop will print the value inside n after the waitgroup is done
	for n:= range c{
		fmt.Println(n)
	}
}
