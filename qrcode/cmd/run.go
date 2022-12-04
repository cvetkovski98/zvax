package cmd

import (
	"log"
	"net"

	"github.com/cvetkovski98/zvax-common/gen/pbqr"
	"github.com/cvetkovski98/zvax/zvax-qrcode/internal/delivery"
	"github.com/cvetkovski98/zvax/zvax-qrcode/internal/service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	runCommand = &cobra.Command{
		Use:   "run",
		Short: "Run QR code microservice",
		Long:  `Run QR code microservice`,
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
	qrService := service.NewQRCodeService()
	qrGrpc := delivery.NewQRServer(qrService)
	server := grpc.NewServer()
	pbqr.RegisterQRServer(server, qrGrpc)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
