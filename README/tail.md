## Creating Errors

There are several types of errors you can create with this package:

### ErrRequired

Signals that a required value is missing.

```go
	err := &errpath.ErrRequired{}
```

### ErrInvalid

Signals that a value is invalid. You can optionally provide valid values and an explanatory message.

```go
	err := &errpath.ErrInvalid[string]{
	    Value:   "invalid_value",
	    Enum:    []string{"valid1", "valid2"},
	    Message: "must be one of the valid values",
	}
```

### ErrField

Represents an error that occurred in a specific field.

```go
	err := &errpath.ErrField{
	    Field: "fieldName",
	    Err:   &errpath.ErrRequired{},
	}
```

### ErrIndex

Represents an error that occurred at a specific index in a slice.

```go
	err := &errpath.ErrIndex{
	    Index: 3,
	    Err:   &errpath.ErrInvalid[int]{Value: 42},
	}
```

### ErrKey

Represents an error that occurred at a specific key in a map.

```go
	err := &errpath.ErrKey{
	    Key: "keyName",
	    Err: &errpath.ErrRequired{},
	}
```

## Error Chaining

Errors can be nested to form detailed error paths. For example:

```go
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
```

This will produce an error message like:

```
	foo.bar["baz"].qux[3].quux ("corge") is invalid
```
