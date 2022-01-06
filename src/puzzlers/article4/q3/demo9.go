package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	// 被“声明并赋值”的变量必须是多个，并且其中至少有一个是新的变量。这时我们才可以说对其中的旧变量进行了重声明。
	// var n int
	// n, err := io.WriteString(os.Stdout, "Hello, everyone!\n") // n 和 error 都声明过之后再使用短变量声明（没有新的变量）会报错

	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n") // 这里对`err`进行了重声明。
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("%d byte(s) were written.\n", n)

	func() {
		// var n = 10 // 不覆盖外部变量
		// n := 10 // 不覆盖外部变量
		n = 10 // 覆盖外部变量
		fmt.Printf("%d byte(s) were written.\n", n)
	}()

	fmt.Printf("%d byte(s) were written.\n", n)

}
