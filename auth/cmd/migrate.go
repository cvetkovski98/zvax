package cmd

import (
	"log"

	"github.com/cvetkovski98/zvax-auth/pkg/config"
	"github.com/cvetkovski98/zvax-auth/pkg/postgresql"
	"github.com/spf13/cobra"
)

var migrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database",
	Long:  `Migrate database`,
	RunE:  migrate,
}

func migrate(cmd *cobra.Command, args []string) error {
	cfg := config.GetConfig()
	db, err := postgresql.NewPgDb(&cfg.Db, &cfg.Pool)
	if err != nil {
		return err
	}
	if err = postgresql.Migrate(cmd.Context(), db); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
