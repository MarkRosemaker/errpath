package errpath_test

import (
	"testing"

	"github.com/MarkRosemaker/errpath"
)

func TestErrInvalid_bool(t *testing.T) {
	errBool := &errpath.ErrInvalid[bool]{}
	if want := `a value (false) is invalid`; errBool.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errBool)
	}

	errBool.Value = true
	if want := `a value (true) is invalid`; errBool.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errBool)
	}

	errBool.Enum = []bool{false}
	if want := `a value (true) is invalid, must be one of: false`; errBool.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errBool)
	}

	errBool.Message = "some reason"
	if want := `a value (true) is invalid: some reason, must be one of: false`; errBool.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errBool)
	}
}

func TestErrInvalid_int(t *testing.T) {
	type num int

	errInt := &errpath.ErrInvalid[num]{}
	if want := `a value (0) is invalid`; errInt.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errInt)
	}

	errInt.Value = 7
	if want := `a value (7) is invalid`; errInt.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errInt)
	}

	errInt.Message = "it is odd"
	if want := `a value (7) is invalid: it is odd`; errInt.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errInt)
	}

	errInt.Enum = []num{2, 4, 6, 8}
	if want := `a value (7) is invalid: it is odd, must be one of: 2, 4, 6, 8`; errInt.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errInt)
	}
}

func TestErrInvalid_float(t *testing.T) {
	type num float64

	errInt := &errpath.ErrInvalid[num]{}
	if want := `a value (0) is invalid`; errInt.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errInt)
	}

	errInt.Value = 7.4
	if want := `a value (7.4) is invalid`; errInt.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errInt)
	}

	errInt.Message = "it is strange"
	if want := `a value (7.4) is invalid: it is strange`; errInt.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errInt)
	}

	errInt.Enum = []num{2.3, 4.1, 6.7, 8.4}
	if want := `a value (7.4) is invalid: it is strange, must be one of: 2.3, 4.1, 6.7, 8.4`; errInt.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errInt)
	}
}

func TestErrInvalid_complex(t *testing.T) {
	errInt := &errpath.ErrInvalid[complex128]{}
	if want := `a value ((0+0i)) is invalid`; errInt.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errInt)
	}

	errInt.Value = 7.4 + 3i
	if want := `a value ((7.4+3i)) is invalid`; errInt.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errInt)
	}

	errInt.Message = "it is strange"
	if want := `a value ((7.4+3i)) is invalid: it is strange`; errInt.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errInt)
	}

	errInt.Enum = []complex128{2.3, 4.1 - 4i, 6.7 + 7i, 8.4 + 11i}
	if want := `a value ((7.4+3i)) is invalid: it is strange, must be one of: (2.3+0i), (4.1-4i), (6.7+7i), (8.4+11i)`; errInt.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errInt)
	}
}

func TestErrInvalid_struct(t *testing.T) {
	type someStruct struct {
		Foo string
		Bar int
	}
	errBool := &errpath.ErrInvalid[someStruct]{}
	if want := `a value is invalid`; errBool.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errBool)
	}

	errBool.Value = someStruct{Foo: "foo", Bar: 3}
	if want := `a value (errpath_test.someStruct{Foo:"foo", Bar:3}) is invalid`; errBool.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errBool)
	}

	errBool.Enum = []someStruct{{Foo: "baz", Bar: 5}}
	if want := `a value (errpath_test.someStruct{Foo:"foo", Bar:3}) is invalid, must be one of: errpath_test.someStruct{Foo:"baz", Bar:5}`; errBool.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errBool)
	}

	errBool.Message = "some reason"
	if want := `a value (errpath_test.someStruct{Foo:"foo", Bar:3}) is invalid: some reason, must be one of: errpath_test.someStruct{Foo:"baz", Bar:5}`; errBool.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errBool)
	}
}

func TestErrInvalid_pointer_to_struct(t *testing.T) {
	type someStruct struct {
		Foo string
		Bar int
	}
	errBool := &errpath.ErrInvalid[*someStruct]{}
	if want := `a value is invalid`; errBool.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errBool)
	}

	errBool.Value = &someStruct{Foo: "foo", Bar: 3}
	if want := `a value (&errpath_test.someStruct{Foo:"foo", Bar:3}) is invalid`; errBool.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errBool)
	}

	errBool.Enum = []*someStruct{{Foo: "baz", Bar: 5}}
	if want := `a value (&errpath_test.someStruct{Foo:"foo", Bar:3}) is invalid, must be one of: &errpath_test.someStruct{Foo:"baz", Bar:5}`; errBool.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errBool)
	}

	errBool.Message = "some reason"
	if want := `a value (&errpath_test.someStruct{Foo:"foo", Bar:3}) is invalid: some reason, must be one of: &errpath_test.someStruct{Foo:"baz", Bar:5}`; errBool.Error() != want {
		t.Fatalf("want: %s, got: %v", want, errBool)
	}
}
