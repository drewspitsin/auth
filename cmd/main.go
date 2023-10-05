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

	"github.com/drewspitsin/auth/internal/config"
	"github.com/drewspitsin/auth/internal/grpc_api"
	desc "github.com/drewspitsin/auth/pkg/user_api_v1"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err.Error())
	}

	greatQuit := make(chan os.Signal, 1)
	signal.Notify(greatQuit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	server := grpc_api.NewUserV1Server()

	desc.RegisterUserV1Server(s, server)

	log.Printf("server listening at %v", lis.Addr())

	go func() {
		if err = s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-greatQuit
	s.GracefulStop()
	fmt.Println(color.YellowString("Auth_api server closed..."))
}
