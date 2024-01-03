package incomingdata

import (
	"fmt"
	"strings"

	cf "github.com/iostrovok/go-convert"
	"github.com/pkg/errors"
)

type Result struct {
	Key   string
	Error error
}

type Incoming struct {
	results []Result
}

// New creates a new instance of the Incoming struct.
// Just a simple constructor.
func New() *Incoming {
	return &Incoming{
		results: []Result{},
	}
}

// First returns the first error as all errors or nil if no error is found.
func (i *Incoming) First() error {
	for k := range i.results {
		if i.results[k].Error != nil {
			return i.results[k].Error
		}
	}

	return nil
}

// All returns the list of results in the Incoming struct.
func (i *Incoming) All() []Result {
	return i.results
}

// Print is a debug function.
// It iterates over the list of results in the Incoming struct and prints each result's Key and Error.
// It returns the Incoming struct itself.
func (i *Incoming) Print() *Incoming {
	s := make([]string, 0)
	for k := range i.results {
		s = append(s, fmt.Sprintf("%s: %+v", i.results[k].Key, i.results[k].Error))
	}

	fmt.Println(strings.Join(s, "\n"))

	return i
}

// Add adds a new check case to the Incoming object.
// It takes in a key(string), value(any type), and typ(Type).
// It returns the Incoming struct itself.
func (i *Incoming) Add(key string, value any, typ Type) *Incoming {
	i.results = append(i.results, Result{
		Key:   key,
		Error: checkAndLogIncomingData(key, value, typ),
	})

	return i
}

func errorf(key string, value any, typ Type, err error) error {
	return errors.Errorf("wrong parameter. Expected: %s. %s => %+v. Convert error: %+v", nameOfType(typ), key, value, err)
}

func checkAndLogIncomingData(key string, value any, typ Type) error {
	switch typ {
	case Int32, Int32P, Int32N:
		a, err := cf.Int32Err(value)
		if err != nil {
			return errorf(key, value, typ, err)
		}

		if (typ == Int32N && a < 0) || (typ == Int32P && a <= 0) {
			return errors.Errorf("%s %d", nameOfType(typ), a)
		}
	case ListInt32N, ListInt32P:
		s, err := cf.ListOfInt32Err(value, true)
		if err != nil {
			return errorf(key, value, typ, err)
		}
		for _, s := range s {
			if s < 0 || (s == 0 && typ == ListInt32P) {
				return errors.Errorf("%s %d", nameOfType(typ), s)
			}
		}
	case Int64, Int64N, Int64P:
		a, err := cf.Int64Err(value)
		if err != nil {
			return errorf(key, value, typ, err)
		}

		if (typ == Int64N && a < 0) || (typ == Int64P && a <= 0) {
			return errors.Errorf("%s %d", nameOfType(typ), a)
		}
	case Float32, Float32N, Float32P:
		a, err := cf.Float32Err(value)
		if err != nil {
			return errorf(key, value, typ, err)
		}

		if (typ == Float32N && a < 0.0) || (typ == Float32P && a <= 0.0) {
			return errorf(key, value, typ, err)
		}
	case Float64, Float64N, Float64P:
		a, err := cf.Float64Err(value)
		if err != nil {
			return errorf(key, value, typ, err)
		}

		if (typ == Float64N && a < 0.0) || (typ == Float64P && a <= 0.0) {
			return errorf(key, value, typ, err)
		}
	case String:
		// Nothing. Any value can be converted to a string (maybe empty).
	case StringP:
		a := cf.String(value)
		if typ == StringP && len(a) == 0 {
			return errorf(key, value, typ, errors.Errorf("empty string"))
		}
	case ListStringP:
		if _, err := cf.ListOfStringsPErr(value, true); err != nil {
			return errorf(key, value, typ, err)
		}
	case ListString:
		if _, err := cf.ListOfStringsErr(value, true); err != nil {
			return errorf(key, value, typ, err)
		}
	case ListStringS:
		if _, err := cf.ListOfStringsStrictPErr(value, true); err != nil {
			return errorf(key, value, typ, err)
		}
	default:
		//
	}

	return nil
}
