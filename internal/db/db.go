package db

import (
	"database/sql"
	"fmt"
	"path/filepath"

	_ "modernc.org/sqlite" // Pure Go SQLite driver
)

func Connect() (*sql.DB, error) {
	dbPath := filepath.Join(".ovinnik", "ovinnik.db")
	conn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka koneksi database: %w", err)
	}
	
	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("gagal melakukan ping ke database: %w", err)
	}

	return conn, nil
}

func InitSchema(conn *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS findings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		check_id TEXT NOT NULL,
		resource TEXT NOT NULL,
		file_path TEXT NOT NULL,
		raw_json TEXT NOT NULL,
		priority INTEGER DEFAULT 0,
		UNIQUE(check_id, resource, file_path)
	);`
	
	_, err := conn.Exec(query)
	if err != nil {
		return fmt.Errorf("gagal membuat schema database: %w", err)
	}
	
	return nil
}