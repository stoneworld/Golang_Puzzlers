package main

import (
	"fmt"
)

func main()  {
	ch := make(chan string)
	go test(ch)
	go test(ch)
	for i := 1; i < 4; i++ {
		_, ok := <-ch               // channel receive
		fmt.Println(ok)
	}
}

func test(ch chan <- string) {
	ch <- "sss12"
}