package main

import (
	"github.com/cvetkovski98/zvax-common/gen/pbslot"
	"github.com/cvetkovski98/zvax-slots/internal/delivery"
	"github.com/cvetkovski98/zvax-slots/internal/repository"
	"github.com/cvetkovski98/zvax-slots/internal/service"
	"github.com/cvetkovski98/zvax-slots/pkg/redis"
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
	rdb, err := redis.NewRedisConn()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	slotRepository := repository.NewRedisSlotRepository(rdb)
	slotService := service.NewSlotServiceImpl(slotRepository)
	slotGrpc := delivery.NewSlotGrpcServerImpl(slotService)
	server := grpc.NewServer()
	pbslot.RegisterSlotGrpcServer(server, slotGrpc)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
