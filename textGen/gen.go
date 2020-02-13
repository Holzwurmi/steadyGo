// Package textGen generates pseudo randomized text from a string source through different methods
package textGen

import (
	"io"
	"strings"
)

type MarkovClusterNode struct {
	State      string
	matrix     transitionMatrix
	nextStates *MarkovClusterNode
}

type transitionMatrix struct {
	probabilties []float64
}

func NewMarkovByWordWithReaderSource(s *io.Reader, degree int) MarkovClusterNode {
	panic("Not implemented")
}

func NewMarkovByRuneWithReaderSource(s *io.Reader, degree int) MarkovClusterNode {
	panic("Not implemented")
}

func NewMarkovByWordWithStringSource(s string, degree int) MarkovClusterNode {
	reader := io.Reader(strings.NewReader(s))
	return NewMarkovByWordWithReaderSource(&reader, degree)
}

func NewMarkovByRuneWithStringSource(s string, degree int) MarkovClusterNode {
	reader := io.Reader(strings.NewReader(s))
	return NewMarkovByRuneWithReaderSource(&reader, degree)
}

// Challenge_Aim: RecursiveData, Probability
