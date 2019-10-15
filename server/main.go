package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	api "github.com/RadioReader/api/handler.go"
)
const(
	PORT = 14586
)


func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		fmt.Printf("Server starting on port %d", PORT)
	}


	s := api.Server{}
	grpcServer := grpc.NewServer()

	// attach the RR service to the server
	api.RegisterRadioReaderServiceServer(grpcServer, & s)

	if err := grpcServer.Serve(lis);
	err != nil {
		log.Fatalf("failed to serve: %s", err)
	} else {
		log.Printf("Server started successfully")
	}


}