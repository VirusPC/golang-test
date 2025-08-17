package main

import (
	"fmt"

	"github.com/VirusPC/golang-test/src/exp"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(exp.Z)
	fmt.Println(exp.ToBe)
}
