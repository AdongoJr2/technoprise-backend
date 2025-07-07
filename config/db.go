package config

import (
	"context"
	"fmt"
	"log"

	"github.com/AdongoJr2/technoprise-backend/ent"
	_ "github.com/lib/pq"
)

// ConnectDB initializes and returns an Ent client connected to PostgreSQL.
// It also runs database migrations.
func ConnectDB(cfg *Config) (*ent.Client, error) {
	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseName,
		cfg.DatabaseSSLMode,
	)

	client, err := ent.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %w", err)
	}

	// Run migration to create the schema.
	// This is suitable for development. For production, consider
	// a dedicated migration tool or a more controlled migration process.
	log.Println("Running database migrations...")
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %w", err)
	}
	log.Println("Database migrations completed successfully.")

	return client, nil
}
