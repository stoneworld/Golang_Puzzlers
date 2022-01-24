package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	// 一旦主 goroutine 中的代码（也就是main函数中的那些代码）执行完毕，当前的 Go 程序就会结束运行。
	// 所以如果没有下面这行代码上面的程序可能就没办法执行了。
	time.Sleep(2 * time.Millisecond)

}
