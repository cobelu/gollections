package apply

import (
	th "gollections/testhelpers"
	"testing"
)

func TestApply(t *testing.T) {

	t.Run("Apply_empty", func(t *testing.T) {
		t.Parallel()

		s := From[int, int]([]int{})
		a := s.
			Apply(func(a int) int { return a + 1 }).
			Do().
			ToSlice()
		th.WantGot(t, 0, len(a), "have same number of elements")
		th.WantGot(t, []int{}, a, "have correct elements")
	})

	t.Run("Apply_single", func(t *testing.T) {
		t.Parallel()

		s := From[int, int]([]int{1, 2, 3})
		a := s.
			Apply(func(a int) int { return a + 1 }).
			Do().
			ToSlice()
		th.WantGot(t, 3, len(a), "have same number of elements")
		th.WantGot(t, []int{2, 3, 4}, a, "have correct elements")
	})

	t.Run("Apply_many", func(t *testing.T) {
		t.Parallel()

		s := From[int, int]([]int{1, 2, 3})
		a := s.
			Apply(func(a int) int { return a + 1 }).
			Apply(func(a int) int { return a + 2 }).
			Do().
			ToSlice()
		th.WantGot(t, len(a), 3, "have same number of elements")
		th.WantGot(t, []int{4, 5, 6}, a, "have correct elements")
	})

}
