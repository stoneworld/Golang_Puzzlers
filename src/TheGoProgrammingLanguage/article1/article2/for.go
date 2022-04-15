package main

import (
	"fmt"
	"os"
)

func main()  {
	for range os.Args[1:] {
		fmt.Println(1111) // receive from channel ch
	}
}