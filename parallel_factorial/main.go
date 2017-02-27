package main

import (
	"fmt"
	"strconv"
	"runtime"
)

func main() {
	cors := runtime.NumCPU();
	runtime.GOMAXPROCS(cors);


	numbers := genNumbers()
	ch := loadNumbers(numbers);
	out:=factorial(ch)
	for n:= range out{
		fmt.Println("out: "+ strconv.Itoa(n))
	}

}
func genNumbers() []int  {
	var numbers = make([]int,30);
	index:=0;
	for i:=0;i<30;i++{
		numbers[i] = i;
		index++
	}
	return  numbers;
}

func loadNumbers(numbers []int) chan int  {
	c:= make(chan int);
	go func() {
		for i:=range numbers{
			c<-i;
		}
		close(c);
	}()
	return  c;
}
func factorial(ch chan int) chan int  {
	c:= make(chan int);
	go func() {
		for number:= range ch{
			total:=1;
			for i:=number ;i>0;i-- {
				total*=i;
			}
			c <-total;
		}
		//push the result into the channel
		close(c)
	}()
	return  c;
}