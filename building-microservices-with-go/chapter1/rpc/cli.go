package main

import (
	"fmt"

	"go-practice/building-microservices-with-go/rpc/client"
	"go-practice/building-microservices-with-go/rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequestAdd(c)
	fmt.Println(reply.Message)

	reply = client.PerformRequestSubtract(c)
	fmt.Println(reply.Message)
}
