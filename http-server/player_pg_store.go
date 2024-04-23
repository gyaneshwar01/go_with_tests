package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresPlayerStore struct {
	db *sql.DB
}

func NewPostgresPlayerStore() (*PostgresPlayerStore, error) {
	connStr := "postgres://myuser:mypassword@localhost/user_score?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return &PostgresPlayerStore{
		db: db,
	}, nil
}

func (s *PostgresPlayerStore) Init() error {
	query := `
		CREATE TABLE players(
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) UNIQUE NOT NULL,
			score INT
		);
	`
	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresPlayerStore) GetPlayerScore(name string) int {
	var score int
	err := s.db.QueryRow("SELECT score FROM players WHERE name=$1", name).Scan(&score)

	if err != nil {
		log.Fatal(err)
	}

	return score
}

func (s *PostgresPlayerStore) RecordWin(name string) {
	_, err := s.db.Exec("INSERT INTO players (name, score) VALUES ($1, 1) ON CONFLICT (name) DO UPDATE SET score = players.score+1", name)

	if err != nil {
		log.Fatal(err)
	}
}
