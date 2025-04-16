package mock

import (
	api "go-markov-api"
)

// Markov is a mock
type Markov struct {
	TrainFn       func(data string) error
	ProbabilityFn func(test string, sequence []string) (float64, error)
	GenerateFn    func(sequence []string) (string, error)
	ChainFn       func() interface{}
}

var _ api.Markov = (*Markov)(nil)

func (m *Markov) Train(data string) error {
	return m.TrainFn(data)
}

func (m *Markov) Probability(test string, sequence []string) (float64, error) {
	return m.ProbabilityFn(test, sequence)
}

func (m *Markov) Generate(sequence []string) (string, error) {
	return m.GenerateFn(sequence)
}

func (m *Markov) Chain() interface{} {
	return m.ChainFn()
}
