package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	api "go-markov-api"
)

type Server struct {
	Markovs map[string]api.Markov
	Port    string
}

var _ api.Server = (*Server)(nil)

func (s *Server) routes() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(s.notFound())

	r.HandleFunc("/", s.service(s.chain())).Methods("GET")
	r.HandleFunc("/train", s.service(s.train())).Methods("POST")
	r.HandleFunc("/probability", s.service(s.probability())).Methods("POST")
	r.HandleFunc("/generate", s.service(s.generate())).Methods("POST")

	http.Handle("/", r)
}

func (s *Server) Start(ctx context.Context) error {
	s.routes()
	srv := &http.Server{Addr: ":" + s.Port, ReadHeaderTimeout: 2 * time.Second}
	go srv.ListenAndServe()

	ctx, _ = signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}
