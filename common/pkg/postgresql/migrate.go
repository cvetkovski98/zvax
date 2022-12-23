package postgresql

import (
	"context"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func Migrate(ctx context.Context, db *bun.DB, migrations *migrate.Migrations) error {
	m := migrate.NewMigrator(db, migrations)
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
