package testhelpers

import (
	"fmt"
	"reflect"
	"testing"
)

// WantGot is a helper function to compare two values and log an error
// if they are not equal. It is a loose BDD implementation.
func WantGot(t *testing.T, want, got any, behavior string) {
	if !reflect.DeepEqual(want, got) {
		t.Fatalf(fmt.Sprintf("%v should %v", t.Name(), behavior))
	}
}
