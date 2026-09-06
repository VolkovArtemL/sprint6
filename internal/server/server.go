package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	log *log.Logger
	srv *http.Server
}

func New(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	h := handlers.New(logger)

	mux.HandleFunc("/", h.MainPage)
	mux.HandleFunc("/upload", h.UploadFile)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		log: logger,
		srv: httpServer,
	}
}

func (s *Server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}