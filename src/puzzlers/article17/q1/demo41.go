package main

import "fmt"

func main() {
	// 示例1。
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	fmt.Println(numbers1)
	fmt.Println()

	// 示例2。
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
	fmt.Println()

	// 示例2。
	numbers4 := []int{1, 2, 3, 4, 5, 6} // 切片
	maxIndex4 := len(numbers4) - 1
	for i, e := range numbers4 {
		// range 进行了浅拷贝
		// 对数组来说，相当于拷贝了一个新的数组进行迭代，修改原数组不会影响被迭代数组。
		// 而对于切片来说，range拷贝出来的切片与原切片底层是同一个数组，因此对原切片的修改也会影响到被迭代切片
		if i == maxIndex4 {
			numbers4[0] += e
		} else {
			numbers4[i+1] += e
		}
	}
	fmt.Println(numbers4)
	fmt.Println()

	// 示例3。
	numbers3 := []int{1, 2, 3, 4, 5, 6}
	maxIndex3 := len(numbers2) - 1
	for i, e := range numbers3 {
		if i == maxIndex3 {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers3)
}
