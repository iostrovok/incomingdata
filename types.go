package incomingdata

type Type int

const (
	None  Type = iota // unknown for a checking type
	Inter             // any / interface{}
	Any               // any / interface{}

	Int32  // simple integer 32
	Int32N // simple non-negative integer 32
	Int32P // simple positive integer 32

	ListInt32N // list of simple non-negative integer 32
	ListInt32P // list of simple positive integer 32

	Int64  // simple integer 64
	Int64N // simple non-negative integer 64
	Int64P // simple positive integer 64

	Float32  // simple float 32
	Float32N // simple non-negative float 32
	Float32P // simple positive float 32

	Float64  // simple float 64
	Float64N // simple non-negative float 64
	Float64P // simple positive float 64

	String  // simple string
	StringP // non-empty string

	// ListString - list has length > 0
	ListString // non-empty list of any values
	// ListStringP - list has at least one non-empty value
	ListStringP // non-empty list containing at least one non-empty value
	// ListStringS - all values in list are not empty
	ListStringS // non-empty list with all non-empty values
)

func nameOfType(t Type) string {
	switch t {
	case Inter, Any:
		return "any"

	case None:
		return "unknown"

	case Int32:
		return "simple integer 32"
	case Int32N:
		return "simple non-negative integer 32"
	case Int32P:
		return "simple positive integer 32"

	case ListInt32P:
		return "non-empty list of simple positive integer 32"
	case ListInt32N:
		return "non-empty list of simple non-negative integer 32"

	case Int64:
		return "simple integer 64"
	case Int64N:
		return "simple non-negative integer 64"
	case Int64P:
		return "simple positive integer 64"

	case Float32:
		return "simple float 32"
	case Float32N:
		return "simple non-negative float 32"
	case Float32P:
		return "simple positive float 32"

	case Float64:
		return "simple float 64"
	case Float64N:
		return "simple non-negative float 64"
	case Float64P:
		return "simple positive float 64"

	case String:
		return "simple string"
	case StringP:
		return "non-empty string"

	case ListString: // ListString - list has length > 0
		return "non-empty list of any values"
	case ListStringP: // ListStringP - list has at least one non-empty value
		return "non-empty list containing at least one non-empty value"
	case ListStringS: // ListStringS - all values in list are not empty
		return "non-empty list with all non-empty values"
	}

	return ""
}
