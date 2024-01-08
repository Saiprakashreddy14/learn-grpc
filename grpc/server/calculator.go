package main

import (
	"context"
	"grpc/proto"
	"io"
	"log"
)

func (s *Server) Add(ctx context.Context, req *proto.CalculationRequest) (*proto.CalculationResponse, error) {
	log.Printf("Add function was invoked with %v", req)
	return &proto.CalculationResponse{
		Result: req.GetA() + req.GetB(),
	}, nil
}

func (s *Server) Subtract(ctx context.Context, req *proto.CalculationRequest) (*proto.CalculationResponse, error) {
	log.Printf("Subtract function was invoked with %v", req)
	return &proto.CalculationResponse{
		Result: req.GetA() - req.GetB(),
	}, nil
}

func (s *Server) Multiply(ctx context.Context, req *proto.CalculationRequest) (*proto.CalculationResponse, error) {
	log.Printf("Multiply function was invoked with %v", req)
	return &proto.CalculationResponse{
		Result: req.GetA() * req.GetB(),
	}, nil
}

func (s *Server) Divide(ctx context.Context, req *proto.CalculationRequest) (*proto.CalculationResponse, error) {
	log.Printf("Divide function was invoked with %v", req)
	return &proto.CalculationResponse{
		Result: req.GetA() / req.GetB(),
	}, nil
}

func (s *Server) PrimeNumberDecomposition(req *proto.PrimeNumberDecompositionRequest, server proto.Calculator_PrimeNumberDecompositionServer) error {
	number := req.Number
	k := 2
	for number > 1 {
		if number%int32(k) == 0 {
			server.Send(&proto.CalculationResponse{Result: number})
			number = number / int32(k)
		} else {
			k = k + 1
		}
	}
	return nil
}

func (s *Server) ComputeAverage(server proto.Calculator_ComputeAverageServer) error {
	sum := int32(0)
	count := 0
	for {
		req, err := server.Recv()
		if err == io.EOF {
			average := float64(sum) / float64(count)
			return server.SendAndClose(&proto.CalculationResponse{
				Result: int32(average),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream : %v", err)
		}
		sum += req.Number
		count++
	}
}

func (s *Server) FindMaximum(server proto.Calculator_FindMaximumServer) error {
	maximum := int32(0)
	for {
		req, err := server.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream : %v", err)
		}
		number := req.Number
		if number > maximum {
			maximum = number
			err = server.Send(&proto.CalculationResponse{Result: maximum})
			if err != nil {
				log.Fatalf("Error while sending data to client : %v", err)
				return err
			}
		}

	}
	return nil
}
