package main

import (
	"fmt"
	"github.com/yugaraxy/samples/logtracer"
)

func a() {
	defer logtracer.Untrace(logtracer.Trace("func a()"))
	fmt.Println("in a")
}

func b() {
	defer logtracer.Untrace(logtracer.Trace("func b"))
	fmt.Println("in b")
	a()
}

func main() {
	b()
}
