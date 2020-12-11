package expectate

import (
	"fmt"
	"strconv"

	"github.com/google/go-cmp/cmp"
)

// Fataler ...
type Fataler interface {
	Fatal(args ...interface{})
}

// Expector ...
type Expector struct {
	t   Fataler
	sub interface{}
}

// ExpectorGenerator ...
type ExpectorGenerator func(subject interface{}) *Expector

// Expect ...
func Expect(t Fataler) ExpectorGenerator {
	expector := new(Expector)
	expector.t = t
	return func(subject interface{}) *Expector {
		expector.sub = subject
		return expector
	}
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

// ToBeEmptyStr ...
func (e Expector) ToBeEmptyStr() {
	str, ok := e.sub.(string)
	if !ok || str != "" {
		e.t.Fatal(e.sub, "is not an empty string")
	}
}

// ToBeZero ...
func (e Expector) ToBeZero() {
	numAsStr := fmt.Sprint(e.sub)
	numAsInt, err := strconv.Atoi(numAsStr)
	if err != nil || numAsInt != 0 {
		e.t.Fatal(e.sub, "is not zero")
	}
}

// ToBeNil ...
func (e Expector) ToBeNil() {
	if e.sub != nil {
		e.t.Fatal(e.sub, "is not nil")
	}
}
