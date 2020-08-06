package main

import (
	"bytes"
	"fmt"
)

/**
 * first GOlang experiments
 */
func main() {
	var myFunc = func(a, b int) int {
		return a + b
	}
	fmt.Printf("hello world %d!\n\n", myFunc(1, 2))

	var array = []int{11, 22, 33, 44}
	for idx, val := range array {
		fmt.Println(idx, val)
	}

	var s = [][]byte{[]byte("hello"), []byte("world")}

	fmt.Printf("%s\n", bytes.Join(s, []byte(", ")))

}
