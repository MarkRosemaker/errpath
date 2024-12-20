package errpath_test

import (
	"testing"

	"github.com/MarkRosemaker/errpath"
)

func TestErrRequired(t *testing.T) {
	err := &errpath.ErrRequired{}
	if want := `a value is required`; err.Error() != want {
		t.Fatalf("want: %s, got: %v", want, err)
	}
}
