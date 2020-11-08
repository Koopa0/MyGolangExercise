package main

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/008-ninja-level-twelve/01/dog"
	dog2 "github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/01/finished/dog"
)

type canine struct {
	name string
	age int
}

func main()  {
	fido := canine{
		name : "Koopa",
		age : dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog2.YearsTwo(20))
}
