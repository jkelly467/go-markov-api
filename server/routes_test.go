package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	api "go-markov-api"
	"go-markov-api/mock"
)

func TestRoutes(t *testing.T) {
	server := new(Server)
	server.routes()

	mockChain := struct {
		Name  string `json:"name"`
		Order int    `json:"order"`
	}{
		"mock chain", 2,
	}

	good := &mock.Markov{
		TrainFn: func(data string) error {
			return nil
		},
		ProbabilityFn: func(test string, sequence []string) (float64, error) {
			return 0.5, nil
		},
		GenerateFn: func(sequence []string) (string, error) {
			return "next", nil
		},
		ChainFn: func() interface{} {
			return mockChain
		},
	}

	bad := &mock.Markov{
		TrainFn: func(data string) error {
			return errors.New("cannot train")
		},
		ProbabilityFn: func(test string, sequence []string) (float64, error) {
			return 0.0, errors.New("cannot calculate")
		},
		GenerateFn: func(sequence []string) (string, error) {
			return "", errors.New("cannot generate")
		},
		ChainFn: func() interface{} {
			return mockChain
		},
	}

	table := []struct {
		markov   api.Markov
		method   string
		path     string
		body     io.Reader
		status   int
		contains string
	}{
		{good,
			"POST",
			"/train",
			strings.NewReader(`{"body": "test"}`),
			200,
			"Success",
		},
		{bad,
			"POST",
			"/train",
			strings.NewReader(`{"body": "test"}`),
			500,
			"cannot train",
		},
		{good,
			"POST",
			"/probability",
			strings.NewReader(`{"test_string": "test", "sequence": ["test", "seq"]}`),
			200,
			"0.5",
		},
		{bad,
			"POST",
			"/probability",
			strings.NewReader(`{"test_string": "test", "sequence": ["test", "seq"]}`),
			500,
			"cannot calculate",
		},
		{good,
			"POST",
			"/generate",
			strings.NewReader(`{"sequence": ["test", "seq"]}`),
			200,
			"next",
		},
		{bad,
			"POST",
			"/generate",
			strings.NewReader(`{"sequence": ["test", "seq"]}`),
			500,
			"cannot generate",
		},
		{good,
			"GET",
			"/",
			nil,
			200,
			"mock chain",
		},
	}

	for _, tt := range table {
		server.Markovs = map[string]api.Markov{
			"two":   tt.markov,
			"three": tt.markov,
		}
		req, _ := http.NewRequest(tt.method, tt.path, tt.body)
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, req)

		assert.Equal(t, tt.status, w.Code, fmt.Sprintf("%v", tt))
		assert.Contains(t, w.Body.String(), tt.contains, fmt.Sprintf("%v", tt))
	}

}
