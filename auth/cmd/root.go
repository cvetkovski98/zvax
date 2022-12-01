package cmd

import (
	"fmt"

	"github.com/cvetkovski98/zvax-auth/internal/config"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Short: "Auth microservice",
	Long:  `Auth microservice`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Auth microservice")
	},
}

func init() {
	cobra.OnInitialize(configure)
	root.AddCommand(runCommand)
	root.AddCommand(migrateCommand)
}

func configure() {
	if err := config.LoadConfig("config.dev.yaml"); err != nil {
		panic(err)
	}
}

func Execute() error {
	return root.Execute()
}
