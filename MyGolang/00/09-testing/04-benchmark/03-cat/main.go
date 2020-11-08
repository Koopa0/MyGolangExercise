package main

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/009-testing/04-benchmark/03-cat/mystr"
	"strings"
)

const s = "Hello"

func main()  {
	xs := strings.Split(s, " ")
	for _,v := range xs{
		fmt.Println(v)
	}
	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n\n", mystr.Join(xs))
}