package main

func main() {
	//pattern - func (chan)  -> async action -> return chan to a new concurrent function
	//multiple stages that run go routines
	ch:= numbersToCalculate(1,2,3)
	out:=sq(ch);
	for n:=range out{
		println(n);
	}
}
func numbersToCalculate(nums ...int) chan int{
	var ch = make(chan int);
	go func() {
	for _,num := range nums{
		ch<-num;
	}
	close(ch);
	}();
	return ch;
}
func sq(in chan int) chan int{
	var ch = make(chan int);
	go func() {
		for num := range in{
			ch<-num * num;
		}
		close(ch);
	}();
	return ch;
}

