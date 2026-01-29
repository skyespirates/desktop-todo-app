// internal/storage/sqlite.go
package storage

import (
	"database/sql"
	"log"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func OpenDB(appName string) (*sql.DB, error) {
	dir, err := AppDataDir(appName)
	if err != nil {
		return nil, err
	}

	dbPath := filepath.Join(dir, "app.db")
	log.Printf("DATABASE PATH: %s", dbPath)

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	// SQLite recommended pragmas
	pragmas := `
	PRAGMA journal_mode = WAL;
	PRAGMA foreign_keys = ON;
	PRAGMA busy_timeout = 5000;
	`
	_, err = db.Exec(pragmas)
	if err != nil {
		return nil, err
	}

	return db, nil
}
