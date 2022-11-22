package main

import (
	"log"
	"net"

	"github.com/cvetkovski98/zvax-common/gen/pbkey"
	"github.com/cvetkovski98/zvax-keys/internal/delivery"
	"github.com/cvetkovski98/zvax-keys/internal/repository"
	"github.com/cvetkovski98/zvax-keys/internal/service"
	"github.com/cvetkovski98/zvax-keys/pkg/postgresql"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening on %v", lis.Addr())
	connPool, err := postgresql.NewPgxConn("postgresql://postgres:changeme@localhost:5432/keys")
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	keyRepository := repository.NewPgKeyRepository(connPool)
	keyService := service.NewKeyServiceImpl(keyRepository)
	keyGrpc := delivery.NewKeyGrpcImpl(keyService)
	server := grpc.NewServer()
	pbkey.RegisterKeyGrpcServer(server, keyGrpc)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
