package main

import (
	"fmt"
)
// using range to handle the waiting for the go routine
func main() {
	c:= make(chan int)

	go func(){
		for i:=0;i<10;i++  {
			c<-i
		}

	}()
// the range will wait until it gets the value of i
	for n:= range c{
		fmt.Println(n)
	}
}
