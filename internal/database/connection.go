package database

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/tursodatabase/go-libsql"
)

func Connect(dbURL, authToken string) (*sql.DB, error) {
	dbName := "local.db"

	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)

	dbPath := filepath.Join(dir, dbName)
	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, dbURL, libsql.WithAuthToken(authToken))

	if err != nil {
		return nil, err
	}
	db := sql.OpenDB(connector)
	return db, nil
}
