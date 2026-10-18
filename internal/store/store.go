package store

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

type Store struct {
	db *sql.DB
}

func Open(path string) (*Store, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Close() error { return s.db.Close() }

func (s *Store) Migrate() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS stories (
			id         TEXT PRIMARY KEY,
			title      TEXT NOT NULL,
			body       TEXT NOT NULL,
			status     TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		);
		CREATE TABLE IF NOT EXISTS tasks (
			id         TEXT PRIMARY KEY,
			story_id   TEXT NOT NULL REFERENCES stories(id),
			title      TEXT NOT NULL,
			body       TEXT NOT NULL,
			status     TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		);
		CREATE TABLE IF NOT EXISTS runs (
			id      TEXT PRIMARY KEY,
			task_id TEXT NOT NULL REFERENCES tasks(id),
			runner  TEXT NOT NULL,
			title   TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS events (
			id      INTEGER PRIMARY KEY AUTOINCREMENT,
			run_id  TEXT NOT NULL REFERENCES runs(id),
		 kind    TEXT NOT NULL,
			payload BLOB,
			ts      DATETIME NOT NULL
		);
	`)
	return err
}
