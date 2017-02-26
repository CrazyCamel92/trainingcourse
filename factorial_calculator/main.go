package main

import (
	"fmt"
	"strconv"
)

func main() {
	num:=4;
	f:=factorial(num)
	fmt.Println("factorial of num: "+ strconv.Itoa(num))
	fmt.Println(<-f)

}
func factorial(num int) chan int  {
	c:= make(chan int);
	go func() {
		total:=1;
		for i:=num ;i>0;i-- {
			total*=i;
	}
		c <-total;
		close(c)
	}()
	return  c;
}