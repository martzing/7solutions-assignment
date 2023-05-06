package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/martzing/7solutions-assignment/3-pie-fire-dire/api"
	"github.com/martzing/7solutions-assignment/3-pie-fire-dire/configs"
	"github.com/martzing/7solutions-assignment/3-pie-fire-dire/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	configs.BootConfig()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *configs.GRPCPort))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	pb.RegisterBeefServiceServer(s, &server.GRPCServer{})
	log.Printf("Serving gRPC on %s:%s\n", *configs.GRPCHost, *configs.GRPCPort)

	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%s:%s", *configs.GRPCHost, *configs.GRPCPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = pb.RegisterBeefServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", *configs.GRPCGateWayPort),
		Handler: gwmux,
	}

	log.Printf("Serving gRPC-Gateway on http://%s:%s\n", *configs.GRPCGateWayHost, *configs.GRPCGateWayPort)
	log.Fatalln(gwServer.ListenAndServe())
}
