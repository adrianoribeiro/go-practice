package client

import (
	"fmt"
	"log"
	"net/rpc"

	"go-practice/building-microservices-with-go/rpc/contract"
)

const port = 1234

func CreateClient() *rpc.Client {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	return client
}

func PerformRequestAdd(client *rpc.Client) contract.CalcResponse {
	args := &contract.CalcRequest{
		Number1:   10,
		Number2:   5,
		Operation: "+",
	}
	var reply contract.CalcResponse

	err := client.Call("CalcHandler.Calc", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}

func PerformRequestSubtract(client *rpc.Client) contract.CalcResponse {
	args := &contract.CalcRequest{
		Number1:   10,
		Number2:   5,
		Operation: "-",
	}
	var reply contract.CalcResponse

	err := client.Call("CalcHandler.Calc", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}
