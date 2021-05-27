package main

import (
	"fmt"
)

func main() {
	s := []interface{}{
		"hello", 11, "dlfk",
	}

	var counter int
	for i := 0; i < len(s); i++ {
		fmt.Println(s[counter])
		counter++
	}
	fmt.Println(len(s))

	fmt.Println(s[1].(int))
}
