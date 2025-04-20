// Package server is the HTTP server with database connection
package server

import (
	"aliserver/storage"
	"encoding/json"
	"net/http"
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
	http.HandleFunc("/top", s.enableCORS(s.handlerTop10()))
	http.HandleFunc("/add", s.enableCORS(s.handlerAddScore()))

	return http.ListenAndServe(":3000", nil)
}

func (s *server) handlerAddScore() handleFunc {
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

		err = s.store.AddScore(score.Name, score.Score)
		if err != nil {
			http.Error(w, "Could not save score"+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *server) handlerTop10() handleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		top10, err := s.store.Top10()
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
