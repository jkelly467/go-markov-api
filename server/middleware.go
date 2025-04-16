package server

import (
	"context"
	"errors"
	"net/http"
)

// define a "context key" type for adding values to the request context
type ctxKey int

const (
	// define a globally unique constant for the "service" context attribute
	ctxService ctxKey = iota
)

func (s *Server) service(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// pull the service off of the query string
		svc := r.URL.Query().Get("service")
		// default to "two"
		if svc == "" {
			svc = "two"
		}

		found := false
		// make sure the requested service is one of the offered services
		for key := range s.Markovs {
			if key == svc {
				found = true
			}
		}

		if !found {
			writeJSON(w, 400, errors.New("bad service"))
		}

		// create a new request context based off the existing context, with our new value added
		ctx := context.WithValue(r.Context(), ctxService, svc)
		// run the endpoint function with the new context
		h(w, r.WithContext(ctx))
	}
}
