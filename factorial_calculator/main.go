package main

import (
	"fmt"
	"strconv"
)

func main() {
	num:=4;
	f:=factorial(num)
	fmt.Println("factorial of num: "+ strconv.Itoa(num))
	//extract a single value from the channel
	fmt.Println(<-f)

}
//mission: calculate the factorial of a number using concurrency
func factorial(num int) chan int  {
	c:= make(chan int);
	go func() {
		total:=1;
		for i:=num ;i>0;i-- {
			total*=i;
	}
		//push the result into the channel
		c <-total;
		close(c)
	}()
	return  c;
}