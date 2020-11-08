package main

import (
	"fmt"
)

func main()  {
	fmt.Println(Greet("Koopa"))
}
func Greet(s string) string  {
	return fmt.Sprint("Hello my dear", s)
}
