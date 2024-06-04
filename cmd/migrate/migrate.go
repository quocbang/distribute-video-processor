package migrate

import (
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/gorm"
)

func MigrateUp(db *gorm.DB, path string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("cannot get sql db: %v", err)
	}

	total, err := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("cannot execute migration: %v", err)
	}

	fmt.Printf("Migrate successfully: %v\n", total)

	return nil
}
