package stream

type Stream[A any] []A

func New[C any]() Stream[C] {
	return make(Stream[C], 0)
}

func (s Stream[A]) ToSlice() []A {
	return s
}
