// Package textGen generates pseudo randomized text from a string source through different methods
package textGen

import (
	"io"
	"strings"
)

const startOfSentenceState = "_^_"

type MarkovCluster struct {
	root   *MarkovClusterNode
	degree int
}

type MarkovClusterNode struct {
	state      string
	nextStates []*NextState
}

type NextState struct {
	ocurrences int
	node       *MarkovClusterNode
}

func (m *MarkovClusterNode) InsertState(occurence string) {
}

func (m *MarkovClusterNode) GetProbabilitiesForNextState() map[string]int {
	return map[string]int{}
}

func NewMarkovByWordWithReaderSource(r *io.Reader, degree int) *MarkovCluster {
	// lastDegreeStates := make([]string, degree-1)
	m := &MarkovCluster{degree: degree, root: &MarkovClusterNode{state: startOfSentenceState, nextStates: []*NextState{}}}

	// for
	// buffer := make([]byte, 100)
	// i, err := r.Read(buffer)

	return m
}

func NewMarkovByRuneWithReaderSource(r *io.Reader, degree int) *MarkovCluster {
	panic("Not implemented")
}

func NewMarkovByWordWithStringSource(s string, degree int) *MarkovCluster {
	reader := io.Reader(strings.NewReader(s))
	return NewMarkovByWordWithReaderSource(&reader, degree)
}

func NewMarkovByRuneWithStringSource(s string, degree int) *MarkovCluster {
	reader := io.Reader(strings.NewReader(s))
	return NewMarkovByRuneWithReaderSource(&reader, degree)
}

// Challenge_Aim: RecursiveData, Probability
