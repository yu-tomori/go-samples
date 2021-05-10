package main

import (
	"fmt"
)

func main() {
	s := new([]string)
	(*s) = append((*s), "hello")
	fmt.Println((*s)[0])
}
