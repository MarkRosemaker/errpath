package errpath_test

import (
	"testing"

	"github.com/MarkRosemaker/errpath"
)

func TestErrKey(t *testing.T) {
	err := &errpath.ErrKey{
		Key: "foo",
		Err: &errpath.ErrKey{
			Key: "bar",
			Err: testErr,
		},
	}
	if want := `["foo"]["bar"]: a test error`; err.Error() != want {
		t.Fatalf("want: %s, got: %v", want, err)
	}
}
