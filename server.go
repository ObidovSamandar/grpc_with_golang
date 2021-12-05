package main

import (
	"log"
	"net"

	"github.com/obidovsamandar/go-crud-with-grpc/controllers"
	"github.com/obidovsamandar/go-crud-with-grpc/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	db.ConnectionDB()

	lis, err := net.Listen("tcp", ":9006")

	if err != nil {
		log.Fatalf("Failed to listen on port 9006: %v", err)
	}

	a := controllers.Server{}

	server := grpc.NewServer()

	controllers.RegisterAddUserServiceServer(server, &a)
	controllers.RegisterGetUserServiceServer(server, &a)
	controllers.RegisterGetAllUserServiceServer(server, &a)
	controllers.RegisterDeleteUserServiceServer(server, &a)
	controllers.RegisterUpdateUserServiceServer(server, &a)

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to server gRPC server over port: 9006 %v", err)
	}

}
