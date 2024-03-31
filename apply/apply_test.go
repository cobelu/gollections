package apply

import (
	th "gollections/testhelpers"
	"testing"
)

func TestApply(t *testing.T) {

	t.Run("Apply", func(t *testing.T) {
		t.Parallel()
		s := From[int, int]([]int{1, 2, 3})
		a := s.Apply(func(a int) int { return a + 1 }).Do().ToSlice()
		th.WantGot(t, len(a), 3, "have same number of elements")
	})

}
