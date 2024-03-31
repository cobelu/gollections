package apply

import "gollections/stream"

// Function application operator

// FIXME: Can only apply from A to B at this time

type Stream[A any, B any] struct {
	fns   []func(A) B
	slice []A
}

func New[A any, B any]() Stream[A, B] {
	return Stream[A, B]{
		fns:   make([]func(A) B, 0),
		slice: make([]A, 0),
	}
}

func From[A any, B any](a []A) Stream[A, B] {
	return Stream[A, B]{
		fns:   make([]func(A) B, 0),
		slice: a,
	}
}

func (s Stream[A, B]) Apply(fn func(A) B) Stream[A, B] {
	s.fns = append(s.fns, fn)
	return s
}

func (s Stream[A, B]) Do() stream.Stream[B] {
	// TODO: Parallelization
	done := make([]B, len(s.slice))
	for i, e := range s.slice {
		for _, fn := range s.fns {
			done[i] = fn(e)
		}
	}
	return done
}
