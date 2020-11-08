package main

import (
	"log"
	"os"
)

func main()  {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//     fmt.Println("err happend", err)
		log.Println("err happend", err)
		//      log.Fatalln(err)
		//      panic(err)
	}
}
