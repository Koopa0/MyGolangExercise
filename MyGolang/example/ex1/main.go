package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn)  {
	defer c.Close()
	for{
		_, err := io.WriteString(c, time.Now().String() + "\n")
		if err != nil{
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}
}
func main()  {
	listener, err:= net.Listen("tcp", "localhost:8080")
	if err != nil{
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil{
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
