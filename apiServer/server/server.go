// Package server is the HTTP server with database connection
package server

import (
	"aliserver/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	prefixAlirobo = ""
	prefixMemory  = "mem"
	prefixInvader = "inv"
)

func New(store storage.Storage) Server {
	return &server{
		store: store,
	}
}

type Server interface {
	Serve() error
}

type server struct {
	store storage.Storage
}

type userScore struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type handleFunc func(w http.ResponseWriter, r *http.Request)

func (s *server) Serve() error {
	fmt.Println("serving on 3000")
	// fs := http.FileServer(http.Dir("."))
	http.HandleFunc("/top", s.enableCORS(s.handlerTop10(prefixAlirobo)))
	http.HandleFunc("/add", s.enableCORS(s.handlerAddScore(prefixAlirobo)))

	http.HandleFunc("/memtop", s.enableCORS(s.handlerTop10(prefixMemory)))
	http.HandleFunc("/memadd", s.enableCORS(s.handlerAddScore(prefixMemory)))

	http.HandleFunc("/invtop", s.enableCORS(s.handlerTop10(prefixInvader)))
	http.HandleFunc("/invadd", s.enableCORS(s.handlerAddScore(prefixInvader)))
	// http.Handle("/", http.StripPrefix("/", fs))

	return http.ListenAndServe(":3000", nil)
}

func (s *server) handlerAddScore(prefix string) handleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "post request expected", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var score userScore
		err := json.NewDecoder(r.Body).Decode(&score)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		switch prefix {
		case prefixAlirobo:
			err = s.store.AddScore(score.Name, score.Score)
		case prefixMemory:
			err = s.store.AddMemScore(score.Name, score.Score)
		case prefixInvader:
			err = s.store.AddInvScore(score.Name, score.Score)
		default:
			http.Error(w, "Not implemented", http.StatusInternalServerError)
			return
		}
		err = s.store.AddScore(score.Name, score.Score)
		if err != nil {
			http.Error(w, "Could not save score"+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *server) handlerTop10(prefix string) handleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var top10 storage.Scores

		switch prefix {
		case prefixAlirobo:
			top10, err = s.store.Top10()
		case prefixMemory:
			top10, err = s.store.TopMem10()
		case prefixInvader:
			top10, err = s.store.TopInv10()
		default:
			http.Error(w, "Not implemented", http.StatusInternalServerError)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusOK)
			return
		}

		jsonBytes, err := json.Marshal(top10)
		if err != nil {
			http.Error(w, err.Error(), http.StatusOK)
			return
		}

		_, err = w.Write(jsonBytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusOK)
			return
		}
	}
}

func (*server) enableCORS(next handleFunc) handleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
