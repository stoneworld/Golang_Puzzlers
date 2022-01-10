package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(nil == m)

	key := "two"
	elem, ok := m["two"]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n",
		key, elem, ok)

	fmt.Printf("The length of nil map: %d%v%T\n",
		len(m),m,m)

	fmt.Printf("Delete the key-element pair by key %q...\n",
		key)
	delete(m, key)

	elem = 2
	fmt.Println("Add a key-element pair to a nil map...")
	m1 := make(map[string]int)
	fmt.Printf("The length of nil map: %d%v%T\n",
		len(m1), m1,m1)
	fmt.Println(nil == m1)

	m1["two"] = elem // 这里会引发panic。
}
