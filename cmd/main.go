package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/drewspitsin/auth/pkg/user_api_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUser_API_V1Server
}

func main() {

	great_quit := make(chan os.Signal, 1)
	signal.Notify(great_quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUser_API_V1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	go func() {
		if err = s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-great_quit
	s.GracefulStop()
	fmt.Println(color.YellowString("Auth_api server closed..."))
}
