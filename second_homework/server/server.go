package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/exp/rand"
)

type Server struct {
	*http.Server
}

func NewServer(addr string) *Server {
	srv := &http.Server{
		Addr: addr,
	}

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		fmt.Fprint(w, "v1.0.0")
	})

	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var json_request string
		err := json.NewDecoder(r.Body).Decode(&json_request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		byte_string, err := base64.StdEncoding.DecodeString(json_request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		outputString := string(byte_string)
		resp := outputString
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/hard-op", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		time.Sleep(time.Duration(rand.Intn(1)+15) * time.Second) // чтобы чаще было завершение процесса через контекст

		status := rand.Intn(1)
		if status == 0 {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Oppeartion successful")
	})

	return &Server{srv}
}

func (s *Server) Start() error {
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
