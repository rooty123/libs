package dbhandler

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DBHandler struct {
	Conn *pg.DB
}

const (
	defaultHost     = "localhost"
	defaultPort     = "5432"
	defaultUser     = "postgres"
	defaultPassword = "postgres"
	defaultDBName   = "postgres"
)

func (db *DBHandler) ConnectPg() error {
	host := getEnv("DB_HOST", defaultHost)
	port := getEnv("DB_PORT", defaultPort)
	user := getEnv("DB_USER", defaultUser)
	password := getEnv("DB_PASSWORD", defaultPassword)
	dbName := getEnv("DB_NAME", defaultDBName)

	addr := fmt.Sprintf("%s:%s", host, port)

	db.Conn = pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: dbName,
	})

	// Test the connection
	_, err := db.Conn.Exec("SELECT 1")
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	return nil
}

// RunMigrations runs migrations from the specified directory using golang-migrate
func (db *DBHandler) RunMigrations() {
	migrationsPath := getEnv("MIGRATIONS_PATH", "file://migrations")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_NAME", "postgres"),
	)

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to create migration instance: %v", err))
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(fmt.Sprintf("Failed to run migrations: %v", err))
	}
	fmt.Println("Migrations ran successfully!")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
