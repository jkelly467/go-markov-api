package gomarkovlib

import (
	api "go-markov-api"
	s "strings"

	"github.com/mb-14/gomarkov"
)

type Markov struct {
	Model gomarkov.Chain
}

var _ api.Markov = (*Markov)(nil)

func (m *Markov) Train(data string) error {
	sanitized := s.Replace(data, "\n", " ", -1)
	m.Model.Add(s.Split(sanitized, " "))

	return nil
}

func (m *Markov) Probability(test string, sequence []string) (float64, error) {
	return m.Model.TransitionProbability(test, sequence)
}

func (m *Markov) Generate(sequence []string) (string, error) {
	return m.Model.Generate(sequence)
}

func (m *Markov) Chain() interface{} {
	return m.Model
}
