package main

import (
	"google.golang.org/grpc"
	"grpc/proto"
	"log"
	"net"
)

type Server struct {
	proto.CalculatorServer
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	log.Printf("Server started on port : %v", lis.Addr().String())

	s := grpc.NewServer()
	proto.RegisterCalculatorServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}

}
