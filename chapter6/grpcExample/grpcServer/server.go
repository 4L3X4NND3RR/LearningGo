package main

import (
	"log"
	"net"

	"github.com/4L3X4NND3RR/chapter6/grpcExample/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to create MoneyTransactionServer.
type server struct {
	
}

// MakeTransaction implements MoneyTransactionServer.MakeTransaction
func (s *server) MakeTransaction(ctx context.Context, in *protofiles.TransactionRequest) (*protofiles.TransactionResponse, error) {
	// Use in.Amount, in.From, in.To and perform transaction logic
	log.Printf("Got request for money Transfer....")
	log.Printf("Amount: %f, From A/c:%s, To A/c:%s", in.Amount, in.From, in.To)
	// Do database logic here....
	return &protofiles.TransactionResponse{Confirmation: true}, nil
}

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protofiles.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
