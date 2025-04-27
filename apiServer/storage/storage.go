package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func New() Storage {
	return &storage{}
}

type Scores []Score

type Score struct {
	Name      string
	Score     int
	CreatedAt time.Time
}

type Storage interface {
	Init() error
	AddScore(name string, score int) error
	Top10() (Scores, error)
	AddMemScore(name string, score int) error
	TopMem10() (Scores, error)
	AddInvScore(name string, score int) error
	TopInv10() (Scores, error)
}

type storage struct {
}

// Init database table if not exists
func (s *storage) Init() error {

	err := s.initAlirobo()
	if err != nil {
		return err
	}

	err = s.initMemory()
	if err != nil {
		return err
	}

	return s.initInvader()
}

func (s *storage) initAlirobo() error {
	db, err := s.getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS user_scores (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		score INTEGER NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return err
	}

	// Check if index exists before creating
	var indexExists bool
	row := db.QueryRow(`SELECT EXISTS (
		SELECT 1 FROM pg_indexes 
		WHERE indexname = 'idx_user_scores_score'
	)`)

	row.Scan(&indexExists)

	if !indexExists {
		_, err = db.Exec(`CREATE INDEX idx_user_scores_score ON user_scores (score)`)
		if err != nil {
			log.Fatalf("Failed to create index: %v", err)
		}
	}

	return nil
}

func (s *storage) initMemory() error {
	db, err := s.getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS mem_user_scores (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		score INTEGER NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return err
	}

	// Check if index exists before creating
	var indexExists bool
	row := db.QueryRow(`SELECT EXISTS (
		SELECT 1 FROM pg_indexes 
		WHERE indexname = 'idx_mem_user_scores_score'
	)`)

	row.Scan(&indexExists)

	if !indexExists {
		_, err = db.Exec(`CREATE INDEX idx_mem_user_scores_score ON mem_user_scores (score)`)
		if err != nil {
			log.Fatalf("Failed to create index: %v", err)
		}
	}

	return nil
}

func (s *storage) initInvader() error {
	db, err := s.getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS inv_user_scores (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		score INTEGER NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return err
	}

	// Check if index exists before creating
	var indexExists bool
	row := db.QueryRow(`SELECT EXISTS (
		SELECT 1 FROM pg_indexes 
		WHERE indexname = 'idx_inv_user_scores_score'
	)`)

	row.Scan(&indexExists)

	if !indexExists {
		_, err = db.Exec(`CREATE INDEX idx_inv_user_scores_score ON inv_user_scores (score)`)
		if err != nil {
			log.Fatalf("Failed to create index: %v", err)
		}
	}

	return nil
}

// AddScore insert new score to db
func (s *storage) AddScore(name string, score int) error {
	db, err := s.getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO user_scores (name, score) values ($1,$2)", name, score)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) Top10() (Scores, error) {
	db, err := s.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT name, score, created_at FROM user_scores order by score desc limit 10")
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores Scores
	for rows.Next() {
		var score Score
		err := rows.Scan(&score.Name, &score.Score, &score.CreatedAt)
		if err != nil {
			return nil, err
		}

		scores = append(scores, score)
	}

	if scores == nil {
		return Scores{}, nil
	}

	return scores, nil
}

// AddScore insert new score to db
func (s *storage) AddMemScore(name string, score int) error {
	db, err := s.getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO mem_user_scores (name, score) values ($1,$2)", name, score)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) TopMem10() (Scores, error) {
	db, err := s.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT name, score, created_at FROM mem_user_scores order by score limit 10")
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores Scores
	for rows.Next() {
		var score Score
		err := rows.Scan(&score.Name, &score.Score, &score.CreatedAt)
		if err != nil {
			return nil, err
		}

		scores = append(scores, score)
	}

	if scores == nil {
		return Scores{}, nil
	}

	return scores, nil
}

// AddScore insert new score to db
func (s *storage) AddInvScore(name string, score int) error {
	db, err := s.getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO inv_user_scores (name, score) values ($1,$2)", name, score)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) TopInv10() (Scores, error) {
	db, err := s.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT name, score, created_at FROM inv_user_scores order by score desc limit 10")
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores Scores
	for rows.Next() {
		var score Score
		err := rows.Scan(&score.Name, &score.Score, &score.CreatedAt)
		if err != nil {
			return nil, err
		}

		scores = append(scores, score)
	}

	if scores == nil {
		return Scores{}, nil
	}

	return scores, nil
}

func (s *storage) getConnection() (*sql.DB, error) {
	var err error
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		return nil, fmt.Errorf("getConnection, unable to connect: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("getConnection, ping database: %w", err)
	}

	return db, nil
}
