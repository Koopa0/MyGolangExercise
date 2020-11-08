package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil{
		//      fmt.Println("err happend", err)
		//      fmt.Println("err happend", err)
		log.Fatalln(err)
		//panic(err)
	}
}
func foo()  {
	fmt.Println("When os.Exit() is called, deferred function don't run")
}
