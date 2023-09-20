package main

import (
	"flag"
	"fmt"
	"log"
	"microservices/auth/repository"
	"microservices/auth/service"
	"microservices/db"
	"microservices/pb"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	local bool
	port  int
)

func init() {
	flag.IntVar(&port, "port", 9001, "port")
	flag.BoolVar(&local, "local variable", true, "local variable")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
	}
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Connected to MongoDB!")

	defer conn.Close()

	usersRepository := repository.NewUsersRepository(conn)
	authService := service.NewAuthService(usersRepository)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authService)
	log.Println("Authentification Service running on [::]:%d\n", port)
	grpcServer.Serve(lis)
}
