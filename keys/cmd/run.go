package cmd

import (
	"log"
	"net"

	"github.com/cvetkovski98/zvax-common/gen/pbkey"
	"github.com/cvetkovski98/zvax-keys/internal/config"
	"github.com/cvetkovski98/zvax-keys/internal/delivery"
	"github.com/cvetkovski98/zvax-keys/internal/repository"
	"github.com/cvetkovski98/zvax-keys/internal/service"
	"github.com/cvetkovski98/zvax-keys/pkg/postgresql"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	runCommand = &cobra.Command{
		Use:   "run",
		Short: "Run keys microservice",
		Long:  `Run keys microservice`,
		Run:   run,
	}
	network string
	address string
)

func init() {
	runCommand.Flags().StringVarP(&network, "network", "n", "tcp", "network to listen on")
	runCommand.Flags().StringVarP(&address, "address", "a", ":50052", "address to listen on")
}

func run(cmd *cobra.Command, args []string) {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s://%s...", network, address)
	cfg := config.GetConfig()
	db, err := postgresql.NewPgDb(&cfg.Db, &cfg.Pool)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	keyRepository := repository.NewPgKeyRepository(db)
	keyService := service.NewKeyServiceImpl(keyRepository)
	keyGrpc := delivery.NewKeyServer(keyService)
	server := grpc.NewServer()
	pbkey.RegisterKeyServer(server, keyGrpc)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
