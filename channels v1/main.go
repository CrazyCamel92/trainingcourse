package main

import (
	"fmt"
	"time"
)
// learning channels v1 un buffered channel
func main() {
	c:= make(chan int)

	go func(){
		for i:=0;i<10;i++  {
			c<-i
		}

	}()

	go func(){
		for i:=0;i<10;i++  {
			fmt.Println(<-c)
		}
	}()
	time.Sleep(time.Second)
}
