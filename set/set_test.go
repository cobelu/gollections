package set

import (
	"testing"

	th "gollections/testhelpers"
)

func TestSet(t *testing.T) {

	s1 := FromSlice([]int{1, 2, 3})
	s2 := FromSlice([]int{3, 2, 1})
	s3 := FromSlice([]int{2, 3, 4})

	t.Run("New", func(t *testing.T) {
		t.Parallel()

		s := New[int]()
		th.WantGot(t, 0, s.Size(), "start empty")
	})

	t.Run("Add_something", func(t *testing.T) {
		t.Parallel()

		s := New[int]().Add(1)
		th.WantGot(t, 1, s.Size(), "add element")
	})

	t.Run("Pop", func(t *testing.T) {
		t.Parallel()

		s := FromSlice([]int{1}).Pop(1)
		th.WantGot(t, 0, s.Size(), "remove element")
	})

	t.Run("Union_same", func(t *testing.T) {
		t.Parallel()

		s := s1.Union(s2)
		th.WantGot(t, 3, s.Size(), "combine two sets")
	})

	t.Run("Union_different", func(t *testing.T) {
		t.Parallel()

		s := s1.Union(s3)
		th.WantGot(t, 4, s.Size(), "combine two sets")
	})

	t.Run("Intersection_same", func(t *testing.T) {
		t.Parallel()

		s := s1.Intersection(s2)
		th.WantGot(t, 3, s.Size(), "have same elements")
	})

	t.Run("Intersection_different", func(t *testing.T) {
		t.Parallel()

		s := s1.Intersection(s3)
		th.WantGot(t, 2, s.Size(), "have two elements in common")
	})

	t.Run("Difference_same", func(t *testing.T) {
		t.Parallel()

		s := s1.Difference(s2)
		th.WantGot(t, 0, s.Size(), "have all elements in common")
	})

	t.Run("Difference_different", func(t *testing.T) {
		t.Parallel()

		s := s1.Difference(s3)
		th.WantGot(t, 1, s.Size(), "have all elements in common")
	})

	t.Run("Size", func(t *testing.T) {
		t.Parallel()

		s := FromSlice([]int{1, 2, 3, 4, 5})
		th.WantGot(t, 5, s.Size(), "have 5 elements")
	})

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()

		th.WantGot(t, New[int]().Empty(), true, "show empty for new set")
	})

	t.Run("SubsetOf_true", func(t *testing.T) {
		t.Parallel()

		s := FromSlice([]int{1, 2})
		th.WantGot(t, s.SubsetOf(s1), true, "be a subset")
	})

	t.Run("SubsetOf_false", func(t *testing.T) {
		t.Parallel()

		s := FromSlice([]int{1, 2, 3, 4})
		th.WantGot(t, s.SubsetOf(s1), false, "not be a subset")
	})

}
