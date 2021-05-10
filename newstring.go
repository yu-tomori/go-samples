package main

import (
	"fmt"
)

func main() {
	str := new(string)
	*str = "hello"
	fmt.Println(*str)
}
