package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func Connect(dbURL, authToken string) (*sql.DB, error) {

	db, err := sql.Open("libsql", strings.Join([]string{dbURL, fmt.Sprintf("authToken=%v", authToken)}, "?"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
