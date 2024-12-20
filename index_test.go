package errpath_test

import (
	"testing"

	"github.com/MarkRosemaker/errpath"
)

func TestErrIndex(t *testing.T) {
	err := &errpath.ErrIndex{
		Index: 3,
		Err: &errpath.ErrIndex{
			Index: 5,
			Err:   testErr,
		},
	}
	if want := `[3][5]: a test error`; err.Error() != want {
		t.Fatalf("want: %s, got: %v", want, err)
	}
}
