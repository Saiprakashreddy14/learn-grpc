package main

import (
	"context"
	"grpc/proto"
	"io"
	"log"
)

func doAdd(client proto.CalculatorClient) {
	log.Printf("Add function was invoked with")
	response, err := client.Add(context.Background(), &proto.CalculationRequest{
		A: 10,
		B: 3,
	})
	if err != nil {
		log.Fatalf("Error while calling Add : %v", err)
	}
	log.Printf("Response from Add : %v", response.Result)
}

func doSub(client proto.CalculatorClient) {
	log.Printf("Add function was invoked with")
	response, err := client.Subtract(context.Background(), &proto.CalculationRequest{
		A: 10,
		B: 3,
	})
	if err != nil {
		log.Fatalf("Error while calling Add : %v", err)
	}
	log.Printf("Response from Add : %v", response.Result)
}

func doMul(client proto.CalculatorClient) {
	log.Printf("Add function was invoked with")
	response, err := client.Multiply(context.Background(), &proto.CalculationRequest{
		A: 10,
		B: 3,
	})
	if err != nil {
		log.Fatalf("Error while calling Add : %v", err)
	}
	log.Printf("Response from Add : %v", response.Result)
}

func doDiv(client proto.CalculatorClient) {
	log.Printf("Add function was invoked with")
	response, err := client.Divide(context.Background(), &proto.CalculationRequest{
		A: 10,
		B: 3,
	})
	if err != nil {
		log.Fatalf("Error while calling Add : %v", err)
	}
	log.Printf("Response from Add : %v", response.Result)
}

func getPrimes(client proto.CalculatorClient) {
	log.Printf("Get Primes was Invoked")

	primeNumberDecomposition, err := client.PrimeNumberDecomposition(context.Background(), &proto.PrimeNumberDecompositionRequest{Number: 270})

	if err != nil {
		log.Fatalf("Error while calling PrimeNumberDecomposition : %v", err)
	}

	for {
		// here server is streaming numbers one by one
		// client is listening to the stream and once the stream is closed by server
		// client breaks the loop
		response, err := primeNumberDecomposition.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving response from PrimeNumberDecomposition : %v", err)
		}
		log.Printf("Response from PrimeNumberDecomposition : %v", response.Result)
	}

}

func getAverage(client proto.CalculatorClient) {
	log.Printf("Get Average was Invoked")

	computeAverage, err := client.ComputeAverage(context.Background())

	if err != nil {
		log.Fatalf("Error while calling ComputeAverage : %v", err)
	}

	numbers := []int32{1, 2, 3, 4}

	// client is streaming numbers to server one by one
	// after sending all numbers, client closes the stream
	// and receives the response from server
	for _, number := range numbers {
		err := computeAverage.Send(&proto.AverageRequest{
			Number: number,
		})
		if err != nil {
			log.Fatalf("Error while sending request to ComputeAverage : %v", err)
		}
	}

	response, err := computeAverage.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from ComputeAverage : %v", err)
	}
	log.Printf("Response from ComputeAverage : %v", response.Result)

}

func doFindMaximum(client proto.CalculatorClient) {
	log.Printf("Find Maximum was Invoked")

	findMaximum, err := client.FindMaximum(context.Background())

	if err != nil {
		log.Fatalf("Error while calling FindMaximum : %v", err)
	}

	numbers := []int32{1, 5, 99, 3, 6, 2, 20}

	waitc := make(chan struct{})

	// client is streaming numbers to server one by one
	// after sending all numbers, client closes the stream
	// and receives the response from server
	go func() {
		for _, number := range numbers {
			err := findMaximum.Send(&proto.MaximumRequest{
				Number: number,
			})
			if err != nil {
				log.Fatalf("Error while sending request to FindMaximum : %v", err)
			}
		}
		findMaximum.CloseSend()
	}()

	// client is listening to the stream from server
	// and prints the maximum number
	go func() {
		for {
			response, err := findMaximum.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving response from FindMaximum : %v", err)
				break
			}
			log.Printf("Response from FindMaximum : %v", response.Result)
		}
		close(waitc)
	}()

	// waiting for above go routines to finish
	<-waitc
}
