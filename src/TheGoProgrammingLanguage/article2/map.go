package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	//months := [...]string{1: "January", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "10", 11: "11", 12: "December"}
	//Q2 := months[4:7]
	//summer := months[6:9]
	//fmt.Println(Q2)     // ["April" "May" "June"]
	//fmt.Printf("%d %d %T", len(summer), cap(summer), months)
	//fmt.Println(summer) // ["June" "July" "August"]



	type Block struct {
		Timestamp int64
		Data []byte
		PrevBlockHash []byte
		Hash []byte
	}

	var res *Block

	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}


	if ages["alice"] != 11 || res.Timestamp == 0 {
		fmt.Printf("%d\n", ages["alice"])
	}

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	list1 := [...]string{1: "January", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "10", 11: "11", 12: "December"}
	fmt.Sprintf("%q", list1)

	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)
	io.WriteString(os.Stdout, s) // 简单起见，忽略一些错误

}

func k(list []string) string { return fmt.Sprintf("%q", list) }
