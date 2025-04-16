package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	api "go-markov-api"
)

func (s *Server) notFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, 404, errors.New("resource not found"))
	}
}

func (s *Server) train() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		svc := r.Context().Value(ctxService).(string)
		var req api.TrainingRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			writeJSON(w, 400, errors.New("invalid request format"))
		}
		fmt.Println(req)

		chain := s.Markovs[svc]
		err = chain.Train(req.Body)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, err)
			return
		}

		writeJSON(w, http.StatusOK, "Success")
	}
}

func (s *Server) probability() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		svc := r.Context().Value(ctxService).(string)
		var req api.ProbabilityRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			writeJSON(w, 400, errors.New("invalid request format"))
		}

		fmt.Println(req)

		chain := s.Markovs[svc]
		prob, err := chain.Probability(req.TestString, req.Sequence)
		fmt.Println(prob)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, err)
			return
		}

		writeJSON(w, http.StatusOK, prob)
	}
}

func (s *Server) generate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		svc := r.Context().Value(ctxService).(string)

		var req api.GenerateRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			writeJSON(w, 400, errors.New("invalid request format"))
		}
		fmt.Println(req)

		chain := s.Markovs[svc]
		next, err := chain.Generate(req.Sequence)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, err)
			return
		}

		writeJSON(w, http.StatusOK, next)
	}
}

func (s *Server) chain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		svc := r.Context().Value(ctxService).(string)
		writeJSON(w, http.StatusOK, s.Markovs[svc].Chain())
	}
}
