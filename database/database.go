package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

// Database provides methods for interacting with the database.
type Database interface {
	InsertFileInfo(filename string, size int64) error
	Close() error
}

// SQLiteDB is a SQLite database implementation of the Database interface.
type SQLiteDB struct {
	db *sql.DB
}

// NewSQLiteDB creates a new SQLite database instance.
func NewSQLiteDB(storagePath string) (*SQLiteDB, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Create table if not exists
	if err := createTables(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return &SQLiteDB{db: db}, nil
}

// InsertFileInfo inserts file information into the database.
func (s *SQLiteDB) InsertFileInfo(filename string, size int64) error {
	stmt, err := s.db.Prepare("INSERT INTO file_info (filename, size) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare insert statement: %w", err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(filename, size); err != nil {
		return fmt.Errorf("failed to insert data into SQLite table: %w", err)
	}

	return nil
}

// Close closes the database connection.
func (s *SQLiteDB) Close() error {
	if err := s.db.Close(); err != nil {
		return fmt.Errorf("failed to close database: %w", err)
	}
	return nil
}

func createTables(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS file_info (
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						filename TEXT,
						size INTEGER
					)`)
	if err != nil {
		return err
	}
	return nil
}
