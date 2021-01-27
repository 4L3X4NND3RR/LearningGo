package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/4L3X4NND3RR/chapter6/serverPush/protofiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port      = ":50051"
	noOfSteps = 3
)

type server struct {
}

// MakeTransaction implements MoneyTransactionServer.MakeTransaction
func (s *server) MakeTransaction(in *protofiles.TransactionRequest, stream protofiles.MoneyTransaction_MakeTransactionServer) error {
	log.Printf("Got request for money transfer....")
	log.Printf("Amount: $%f, From A/c: %s, To A/c: %s", in.Amount, in.From, in.To)
	// Send streams here
	for i := 0; i < noOfSteps; i++ {
		time.Sleep(time.Second * 2)
		// Once task is done, send the successful message back to the client
		if err := stream.Send(&protofiles.TransactionResponse{Status: "good", Step: int32(i), Description: fmt.Sprintf("Performing step %d", int32(i))}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, "status", err)
		}
	}
	log.Printf("Successfully transferred amount $%v from %v to %v", in.Amount, in.From, in.To)
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// Create a new GRPC Server and register
	s := grpc.NewServer()
	protofiles.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
