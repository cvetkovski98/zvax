package postgresql

import (
	"context"
	"log"

	"github.com/cvetkovski98/zvax-keys/internal/model/migrations"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func Migrate(ctx context.Context, db *bun.DB) error {
	m := migrate.NewMigrator(db, migrations.Migrations)
	if err := m.Init(ctx); err != nil {
		return err
	}
	if migrations, err := m.Migrate(ctx); err != nil {
		return err
	} else {
		log.Println("Migrations:", migrations)
	}
	return nil
}
