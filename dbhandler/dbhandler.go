package dbhandler

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
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

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
