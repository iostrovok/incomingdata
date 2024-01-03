package incomingdata

import (
	"fmt"
	"math"
	"testing"
)

func TestInt32(t *testing.T) {
	k := Int32

	err := New().
		Add("k", 12, k).
		Add("k", -12, k).
		Add("k", 0, k).
		Add("k", "10", k).
		First()

	Nil(t, err)

	err = New().Add("k", math.MinInt64, k).First()
	NotNil(t, err)

	err = New().Add("k", math.MaxInt64, k).First()
	NotNil(t, err)
}

func TestInt32N(t *testing.T) {
	k := Int32N

	err := New().
		Add("k", 12, k).
		Add("k", 0, k).
		Add("k", "10", k).
		First()

	Nil(t, err)

	err = New().Add("k", -1, k).First()
	NotNil(t, err)

	err = New().Add("k", math.MinInt64, k).First()
	NotNil(t, err)

	err = New().Add("k", math.MaxInt64, k).First()
	NotNil(t, err)
}

func TestInt32P(t *testing.T) {
	k := Int32P

	err := New().
		Add("k", 12, k).
		Add("k", "10", k).
		First()

	Nil(t, err)

	err = New().Add("k", 0, k).First()
	NotNil(t, err)

	err = New().Add("k", -1, k).First()
	NotNil(t, err)

	err = New().Add("k", math.MinInt64, k).First()
	NotNil(t, err)

	err = New().Add("k", math.MaxInt64, k).First()
	NotNil(t, err)
}

func TestInt64(t *testing.T) {
	k := Int64

	err := New().
		Add("k", 12, k).
		Add("k", -12, k).
		Add("k", 0, k).
		Add("k", "10", k).
		First()

	Nil(t, err)
}

func TestInt64N(t *testing.T) {
	k := Int64N

	err := New().
		Add("k", 12, k).
		Add("k", 0, k).
		Add("k", "10", k).
		First()

	Nil(t, err)

	err = New().Add("k", -1, k).First()
	NotNil(t, err)
}

func TestInt64P(t *testing.T) {
	k := Int64P

	err := New().
		Add("k", 12, k).
		Add("k", "10", k).
		First()

	Nil(t, err)

	err = New().Add("k", 0, k).First()
	NotNil(t, err)

	err = New().Add("k", -1, k).First()
	NotNil(t, err)
}

func TestFloat32(t *testing.T) {
	k := Float32

	err := New().
		Add("k", 12, k).
		Add("k", -12, k).
		Add("k", 0, k).
		Add("k", "10", k).
		First()

	Nil(t, err)

	err = New().Add("k", -1*math.MaxFloat64, k).First()
	NotNil(t, err)

	err = New().Add("k", math.MaxFloat64, k).First()
	NotNil(t, err)
}

func TestFloat32N(t *testing.T) {
	k := Float32N

	err := New().
		Add("k", 12, k).
		Add("k", 0, k).
		Add("k", "10", k).
		First()

	Nil(t, err)

	err = New().Add("k", -1, k).First()
	NotNil(t, err)

	err = New().Add("k", -1*math.MaxFloat64, k).First()
	NotNil(t, err)

	err = New().Add("k", math.MaxFloat64, k).First()
	NotNil(t, err)
}

func TestFloat32P(t *testing.T) {
	k := Float32P

	err := New().
		Add("k", 12, k).
		Add("k", "10", k).
		First()

	Nil(t, err)

	err = New().Add("k", 0, k).First()
	NotNil(t, err)

	err = New().Add("k", -1, k).First()
	NotNil(t, err)

	err = New().Add("k", -1*math.MaxFloat64, k).First()
	NotNil(t, err)

	err = New().Add("k", math.MaxFloat64, k).First()
	NotNil(t, err)
}

func TestFloat64(t *testing.T) {
	k := Float64

	err := New().
		Add("Float64-1", 12, k).
		Add("Float64-2", -12, k).
		Add("Float64-3", 0, k).
		Add("Float64-4", "10", k).
		First()

	Nil(t, err)
}

func TestFloat64N(t *testing.T) {
	k := Float64N

	err := New().
		Add("Float64N-1", 12, k).
		Add("Float64N-2", 0, k).
		Add("Float64N-3", "10", k).
		First()

	Nil(t, err)

	err = New().Add("Float64N-4", -1, k).First()
	NotNil(t, err)
}

func TestFloat64P(t *testing.T) {
	k := Float64P

	err := New().
		Add("Float64P-1", 12, k).
		Add("Float64P-2", "10", k).
		First()

	Nil(t, err)

	err = New().Add("Float64P-3", 0, k).First()
	NotNil(t, err)

	err = New().Add("Float64P-4", -1, k).First()
	NotNil(t, err)
}

// StringP - non-empty string
func TestStringP(t *testing.T) {
	k := StringP

	err := New().
		Add("StringP-1", 12, k).
		Add("StringP-2", -12, k).
		Add("StringP-3", 0, k).
		Add("StringP-4", "10", k).
		First()

	Nil(t, err)

	err = New().Add("StringP-5", "", k).First()
	NotNil(t, err)

	err = New().Add("StringP-6", nil, k).First()
	NotNil(t, err)
}

// String - simple string
func TestString(t *testing.T) {
	k := String

	err := New().
		Add("String-1", 12, k).
		Add("String-2", -12, k).
		Add("String-3", 0, k).
		Add("String-4", "10", k).
		Add("String-5", "", k).
		Add("String-6", nil, k).
		First()

	Nil(t, err)
}

// ListStringP - list has at least one non-empty value, non-empty list containing at least one non-empty value
func TestListStringP(t *testing.T) {
	k := ListStringP

	err := New().
		Add("ListStringP-1", []any{"12"}, k).
		Add("ListStringP-2", []any{"-12"}, k).
		Add("ListStringP-3", []any{"0"}, k).
		Add("ListStringP-3", []any{"0", ""}, k).
		Add("ListStringP-4", []any{"2", "1"}, k).
		First()

	Nil(t, err)

	err = New().Add("ListStringP-5", []any{}, k).First()
	NotNil(t, err)
	//
	err = New().Add("ListStringP-6", []any{nil}, k).First()
	NotNil(t, err)
	//
	err = New().Add("ListStringP-7", []any{""}, k).First()
	NotNil(t, err)

	err = New().Add("ListStringP-8", []any{"1", "", nil}, k).Print().First()
	NotNil(t, err)
}

// ListString - list has length > 0, non-empty list of any values
func TestListString(t *testing.T) {
	k := ListString

	err := New().
		Add("ListString-1", []any{"12"}, k).
		Add("ListString-2", []any{"-12"}, k).
		Add("ListString-3", []any{"0"}, k).
		Add("ListString-4", []any{"2", "1"}, k).
		Add("ListString-5", []any{""}, k).
		Add("ListString-7", []any{false}, k).
		First()

	Nil(t, err)

	err = New().Add("ListString-8", []any{}, k).Print().First()
	NotNil(t, err)

	err = New().Add("ListString-9", []string{}, k).Print().First()
	NotNil(t, err)
}

// ListStringS - all values in list are not empty, non-empty list with all non-empty values
func TestListStringS(t *testing.T) {
	k := ListStringS

	err := New().
		Add("ListString-1", []any{"12"}, k).
		Add("ListString-2", []any{"-12"}, k).
		Add("ListString-3", []any{"0"}, k).
		Add("ListString-4", []any{"2", "1"}, k).
		First()

	Nil(t, err)

	err = New().Add("ListString-8", []any{}, k).Print().First()
	NotNil(t, err)

	err = New().Add("ListString-9", []string{}, k).Print().First()
	NotNil(t, err)

	err = New().Add("ListString-9", []any{"2", "1", ""}, k).Print().First()
	NotNil(t, err)
}

func Fatal(t *testing.T, err error) {
	if err != nil {
		fmt.Printf("%+v\n", err)
		t.Fatal(err)
	}
}

func NotNil(t *testing.T, err error) {
	if err == nil {
		Fatal(t, fmt.Errorf("expected not nil errors"))
	}
}

func Nil(t *testing.T, err error) {
	if err != nil {
		Fatal(t, err)
	}
}
