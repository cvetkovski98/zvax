package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Short: "QR code microservice",
	Long:  `QR code microservice`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("QR code microservice")
	},
}

func init() {
	root.AddCommand(runCommand)
}

func Execute() error {
	return root.Execute()
}
