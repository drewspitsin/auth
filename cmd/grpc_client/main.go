package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/drewspitsin/auth/internal/model"
	descAccess "github.com/drewspitsin/auth/pkg/access_v1"
)

var accessToken = flag.String("a", "", "access token")

const (
	servicePort = 50051
	authPrefix  = "Bearer"
	authHeader  = "Authorization"
)

func main() {
	flag.Parse()

	ctx := context.Background()
	md := metadata.New(map[string]string{authHeader: authPrefix + " " + *accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.Dial(
		fmt.Sprintf(":%d", servicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial GRPC client: %v", err)
	}

	cl := descAccess.NewAccessV1Client(conn)

	_, err = cl.Check(ctx, &descAccess.CheckRequest{
		EndpointAddress: model.ExamplePath,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Access granted")
}