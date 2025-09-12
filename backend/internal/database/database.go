package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"goapi/internal/config"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// DB wraps the sql.DB with additional methods
type DB struct {
	*sql.DB
}

// NewDatabase creates a new database connection
func NewDatabase(cfg *config.Config) (*DB, error) {
	// Open database connection
	db, err := sql.Open("postgres", cfg.GetDatabaseURL())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully!")

	// Create tables
	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return &DB{db}, nil
}

// createTables creates all necessary tables
func createTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	// Create trigger to update updated_at timestamp
	triggerQuery := `
	CREATE OR REPLACE FUNCTION update_updated_at_column()
	RETURNS TRIGGER AS $$
	BEGIN
		NEW.updated_at = CURRENT_TIMESTAMP;
		RETURN NEW;
	END;
	$$ language 'plpgsql';

	DROP TRIGGER IF EXISTS update_users_updated_at ON users;
	CREATE TRIGGER update_users_updated_at
		BEFORE UPDATE ON users
		FOR EACH ROW
		EXECUTE FUNCTION update_updated_at_column();
	`
	
	_, err = db.Exec(triggerQuery)
	if err != nil {
		return fmt.Errorf("failed to create updated_at trigger: %w", err)
	}

	log.Println("Database tables created successfully!")
	return nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.DB.Close()
}

// IsUniqueConstraintError checks if the error is a unique constraint violation
func IsUniqueConstraintError(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok {
		return pqErr.Code == "23505" // unique_violation
	}
	return false
}

// IsNoRowsError checks if the error is a no rows error
func IsNoRowsError(err error) bool {
	return err == sql.ErrNoRows
}
