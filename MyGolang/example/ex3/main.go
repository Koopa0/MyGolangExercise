package main

import (
"fmt"
)

func main() {
	fmt.Println(passable(1, increment, increment, increment, increment)) // 4
}


func increment(a int) int {
	a += 1

	return a
}


func passable(a int, funcs ...func(a int) int) int {
	for _, f := range funcs {
		a = f(a)
	}

	return a
}