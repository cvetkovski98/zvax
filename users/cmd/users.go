package main

import (
	"github.com/cvetkovski98/zvax-common/gen/pbuser"
	"github.com/cvetkovski98/zvax-users/internal/users/delivery"
	"github.com/cvetkovski98/zvax-users/internal/users/repository"
	"github.com/cvetkovski98/zvax-users/internal/users/service"
	"github.com/cvetkovski98/zvax-users/pkg/postgresql"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening on %v", lis.Addr())
	connPool, err := postgresql.NewPgxConn("postgresql://postgres:changeme@localhost:5431/users")
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	userRepository := repository.NewPgUserRepository(connPool)
	userService := service.NewUserServiceImpl(userRepository)
	userGrpc := delivery.NewUserGrpcImpl(userService)
	server := grpc.NewServer()
	pbuser.RegisterUserGrpcServer(server, userGrpc)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
