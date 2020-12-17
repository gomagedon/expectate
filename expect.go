package expectate

import (
	"fmt"
	"strconv"

	"github.com/gomagedon/expectate/check"
	"github.com/google/go-cmp/cmp"
)

// Expect ...
func Expect(t Fataler) ExpectorGenerator {
	expector := new(Expector)
	expector.t = t
	return func(subject interface{}) *Expector {
		expector.sub = subject
		return expector
	}
}

// Fataler ...
type Fataler interface {
	Fatal(args ...interface{})
}

// ExpectorGenerator ...
type ExpectorGenerator func(subject interface{}) *Expector

// Expector ...
type Expector struct {
	t   Fataler
	sub interface{}
}

// ToBe ...
func (e Expector) ToBe(expected interface{}) {
	if !check.Is(e.sub, expected) {
		e.t.Fatal(e.sub, "is not", expected)
	}
}

// ToEqual ...
func (e Expector) ToEqual(expected interface{}) {
	diff := cmp.Diff(expected, e.sub)

	if diff != "" {
		e.t.Fatal(diff)
	}
}

// ToBeEmptyStr ...
func (e Expector) ToBeEmptyStr() {
	str, ok := e.sub.(string)
	if !ok || str != "" {
		e.t.Fatal(e.sub, "is not an empty string")
	}
}

// ToBeZero ...
func (e Expector) ToBeZero() {
	switch e.sub.(type) {
	case string:
		e.t.Fatal(e.sub, "is not zero")
	case fmt.Stringer:
		e.t.Fatal(e.sub, "is not zero")
	default:
		numAsStr := fmt.Sprint(e.sub)
		numAsInt, err := strconv.Atoi(numAsStr)
		if err != nil || numAsInt != 0 {
			e.t.Fatal(e.sub, "is not zero")
		}
	}
}

// ToBeNil ...
func (e Expector) ToBeNil() {
	switch e.sub.(type) {
	case string:
		e.t.Fatal(e.sub, "is not nil")
	case fmt.Stringer:
		e.t.Fatal(e.sub, "is not nil")
	default:
		if fmt.Sprint(e.sub) != "<nil>" {
			e.t.Fatal(e.sub, "is not nil")
		}
	}
}

// NotToBe ...
func (e Expector) NotToBe(expected interface{}) {
	if check.Is(e.sub, expected) {
		e.t.Fatal(e.sub, "is", expected)
	}
}

// NotToEqual ...
func (e Expector) NotToEqual(expected interface{}) {
	if cmp.Equal(e.sub, expected) {
		e.t.Fatal(e.sub, "equals", expected)
	}
}
