package main

import (
	"fmt"
)

func main() {
	m := new(map[string]string)
	m = &map[string]string{} // この行がないとエラーになる
	(*m)["hello"] = "world"
	fmt.Println((*m)["hello"])
}
