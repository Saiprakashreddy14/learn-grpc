package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/proto"
	"log"
)

var addr string = "localhost:50051"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	log.Printf("Connecting to %v", addr)

	if err != nil {
		log.Printf("Failed to connect : %v", err)
		panic(err)
	}

	calculatorClient := proto.NewCalculatorClient(conn)

	// unary calls , similar to rest api calls
	doAdd(calculatorClient)
	doSub(calculatorClient)
	doMul(calculatorClient)
	doDiv(calculatorClient)

	// server streaming calls
	getPrimes(calculatorClient)

	// client streaming calls
	getAverage(calculatorClient)

	// bidirectional streaming calls
	doFindMaximum(calculatorClient)

	defer conn.Close()

}
