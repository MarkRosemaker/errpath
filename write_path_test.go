package errpath_test

import (
	"errors"
	"testing"

	"github.com/MarkRosemaker/errpath"
)

var testErr = errors.New("a test error")

func TestPath(t *testing.T) {
	err := &errpath.ErrField{
		Field: "foo",
		Err: &errpath.ErrField{
			Field: "bar",
			Err: &errpath.ErrKey{
				Key: "baz",
				Err: &errpath.ErrField{
					Field: "qux",
					Err: &errpath.ErrIndex{
						Index: 3,
						Err: &errpath.ErrField{
							Field: "quux",
							Err: &errpath.ErrInvalid[string]{
								Value: "corge",
							},
						},
					},
				},
			},
		},
	}
	if want := `foo.bar["baz"].qux[3].quux ("corge") is invalid`; err.Error() != want {
		t.Fatalf("want: %s, got: %s", want, err.Error())
	}
}

func TestPath_join(t *testing.T) {
	err := &errpath.ErrField{
		Field: "foo",
		Err: &errpath.ErrKey{
			Key: "bar",
			Err: errors.Join(
				&errpath.ErrIndex{
					Index: 3,
					Err: &errpath.ErrField{
						Field: "name",
						Err: &errpath.ErrInvalid[string]{
							Value:   "corge",
							Message: "duplicate name",
						},
					},
				},
				&errpath.ErrIndex{
					Index: 5,
					Err: &errpath.ErrField{
						Field: "name",
						Err: &errpath.ErrInvalid[string]{
							Value:   "corge",
							Message: "duplicate name",
						},
					},
				},
			),
		},
	}
	if want := `foo["bar"][3].name ("corge") is invalid: duplicate name
foo["bar"][5].name ("corge") is invalid: duplicate name`; err.Error() != want {
		t.Fatalf("want: %s, got: %s", want, err.Error())
	}
}
