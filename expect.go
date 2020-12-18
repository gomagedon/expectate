package expectate

import (
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
	if e.sub != expected {
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

// NotToBe ...
func (e Expector) NotToBe(expected interface{}) {
	if e.sub == expected {
		e.t.Fatal(e.sub, "is", expected)
	}
}

// NotToEqual ...
func (e Expector) NotToEqual(expected interface{}) {
	if cmp.Equal(e.sub, expected) {
		e.t.Fatal(e.sub, "equals", expected)
	}
}
