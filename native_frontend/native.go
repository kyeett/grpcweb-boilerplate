package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/kyeett/grpcweb-boilerplate/proto/server"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:10000"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("conn yo %s", err)
	}
	defer conn.Close()

	c := pb.NewBackendClient(conn)
	resp, err := c.NewPlayer(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not playnew %s", err)
	}
	fmt.Println("Received", resp.GetID())
}
