package main

import "fmt"

func main() {
	c:= increment();
	sum:=pull(c);
	// the result inside sum ill contain the sum of the numbers 0-9
	for n:= range sum {
		fmt.Print(n); // expected result is (45 0+1+2+3+4+5+6+7+8+9)
	}
}
func increment() chan int{
	out:= make(chan int)
	//creating a new channel and adding 10 values to it.
	go func(){
		for i:=0;i<10;i++  {
			out <-i
		}
		//closing the channel
		close(out)
	}()
	return  out;
}
func pull(c chan int) chan int {
	out:=make(chan int);
	// getting the numbers from 0-9 inside the channel c and adding them to a local int variable sum.
	go func() {
		var sum int;
		for n:= range c{
			sum+=n;
		}
		// the result will pass to the new channel out and return back to main
		out <- sum
		close(out)
	}()
	return  out;
}