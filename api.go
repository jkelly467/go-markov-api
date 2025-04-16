package api

import (
	"context"
)

type Markov interface {
	Train(data string) error
	Probability(test string, sequence []string) (float64, error)
	Generate(sequence []string) (string, error)
	Chain() interface{}
}

type Server interface {
	Start(context.Context) error
}

type TrainingRequest struct {
	Body string `json:"body"`
}

type ProbabilityRequest struct {
	TestString string   `json:"test_string"`
	Sequence   []string `json:"sequence"`
}

type GenerateRequest struct {
	Sequence []string `json:"sequence"`
}
