package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	prefixAlirobo = ""
	prefixMemory  = "mem"
	prefixInvader = "inv"
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

	err := s.initTable(prefixAlirobo)
	if err != nil {
		return err
	}

	err = s.initTable(prefixMemory)
	if err != nil {
		return err
	}

	return s.initTable(prefixInvader)
}

// AddScore insert new score to db
func (s *storage) AddScore(name string, score int) error {
	return s.addScore(name, prefixAlirobo, score)
}

func (s *storage) Top10() (Scores, error) {
	return s.top10(prefixAlirobo, true)
}

// AddScore insert new score to db
func (s *storage) AddMemScore(name string, score int) error {
	return s.addScore(name, prefixMemory, score)
}

func (s *storage) TopMem10() (Scores, error) {
	return s.top10(prefixMemory, true)
}

// AddScore insert new score to db
func (s *storage) AddInvScore(name string, score int) error {
	return s.addScore(name, prefixInvader, score)
}

func (s *storage) TopInv10() (Scores, error) {
	return s.top10(prefixInvader, true)
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

func (s *storage) initTable(prefix string) error {
	db, err := s.getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS ` + prefix + `user_scores (
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
		WHERE indexname = 'idx_` + prefix + `user_scores_score'
	)`)

	row.Scan(&indexExists)

	if !indexExists {
		_, err = db.Exec(`CREATE INDEX idx_` + prefix + `user_scores_score ON user_scores (score)`)
		if err != nil {
			log.Fatalf("Failed to create index: %v", err)
		}
	}

	return nil
}

// AddScore insert new score to db
func (s *storage) addScore(name, prefix string, score int) error {
	db, err := s.getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO "+prefix+"user_scores (name, score) values ($1,$2)", name, score)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) top10(prefix string, desc bool) (Scores, error) {
	db, err := s.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	descParam := ""
	if desc {
		descParam = " desc"
	}

	sql := "SELECT name, score, created_at FROM " + prefix + "user_scores order by score" + descParam + " limit 10"
	stmt, err := db.Prepare(sql)
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
