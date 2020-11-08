package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {
	f, err :=os.Create("names.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer  f.Close()

	r := strings.NewReader("Koopa")

	io.Copy(f, r)
}
