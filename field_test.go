package errpath_test

import (
	"testing"

	"github.com/MarkRosemaker/errpath"
)

func TestErrField(t *testing.T) {
	err := &errpath.ErrField{
		Field: "foo",
		Err: &errpath.ErrField{
			Field: "bar",
			Err:   testErr,
		},
	}
	if want := `foo.bar: a test error`; err.Error() != want {
		t.Fatalf("want: %s, got: %v", want, err)
	}
}
