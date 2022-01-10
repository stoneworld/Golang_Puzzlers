package main

import "fmt"

var block = "package"

func main() {
	fmt.Printf("The block is %s.\n", block)
	var block = 1 // 声明重名的变量是无法通过编译的，用短变量声明对已有变量进行重声明除外
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
}
