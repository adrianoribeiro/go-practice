package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"go-practice/building-microservices-with-go/rpc/contract"
)

const port = 1234

func main() {
	log.Printf("Server starting on port %v\n", port)
	StartServer()
}

func StartServer() {
	calcHandler := &CalcHandler{}
	rpc.Register(calcHandler)

	//Creating a socket using the given protocol and binding it to the IP address and port.
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		//The Fatal functions call os.Exit(1) after writing the log message
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", port))
	}
	defer l.Close()

	for {
		// Accept waits for and returns the next connection to the listener.
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}

type CalcHandler struct{}

func (h *CalcHandler) Calc(args *contract.CalcRequest, reply *contract.CalcResponse) error {

	var result string
	switch args.Operation {

	case "+":
		calc := args.Number1 + args.Number2
		result = fmt.Sprintf("%d %s %d = %d", args.Number1, args.Operation, args.Number2, calc)
	case "-":
		calc := args.Number1 - args.Number2
		result = fmt.Sprintf("%d %s %d = %d", args.Number1, args.Operation, args.Number2, calc)
	default:
		result = "Undefined operation"
	}

	reply.Message = result
	return nil
}
