package main

func main() {
	ch := make(chan int) // ch has type 'chan int'
	ch <- 1
	<-ch // a receive statement; result is discarded
}
