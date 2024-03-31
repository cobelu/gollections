package set

import "maps"

type Set[C comparable] map[C]struct{}

func New[C comparable]() Set[C] {
	return make(Set[C], 0)
}

func FromSlice[C comparable](c []C) Set[C] {
	return make(Set[C], len(c)).AddAll(c)
}

func (s Set[C]) Add(c C) Set[C] {
	s[c] = struct{}{}
	return s
}

func (s Set[C]) AddAll(c []C) Set[C] {
	// TODO: Make this more efficient
	for _, v := range c {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[C]) Pop(c C) Set[C] {
	delete(s, c)
	return s
}

func (s Set[C]) Union(other Set[C]) Set[C] {
	newSet := New[C]()
	maps.Copy(newSet, s)
	maps.Copy(newSet, other)
	return s
}

func (s Set[C]) Intersection(other Set[C]) Set[C] {
	// TODO: Iterate over the smaller set
	newSet := New[C]()
	for k := range s {
		if _, ok := other[k]; ok {
			newSet[k] = struct{}{}
		}
	}
	return newSet
}

func (s Set[C]) Difference(other Set[C]) Set[C] {
	newSet := New[C]()
	for k := range s {
		if _, ok := other[k]; !ok {
			newSet[k] = struct{}{}
		}
	}
	return newSet
}

func (s Set[C]) Contains(c C) bool {
	_, ok := s[c]
	return ok
}

func (s Set[C]) Size() int {
	return len(s)
}

func (s Set[C]) Empty() bool {
	return s.Size() == 0
}

// TODO: Implement this
func (s Set[C]) SubsetOf(other Set[C]) bool {
	for k := range s {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}

func (s Set[C]) ToSlice() []C {
	c := make([]C, s.Size())
	for k := range s {
		c = append(c, k)
	}
	return c
}
