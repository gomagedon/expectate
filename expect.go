package expectate

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// Fataler is satisfied by *testing.T but could be something else if needed
type Fataler interface {
	Fatalf(format string, args ...interface{})
}

// Expect constructs an ExpectorFunc object
func Expect(t Fataler) ExpectorFunc {
	expector := new(Expector)
	expector.t = t
	return func(subject interface{}) *Expector {
		expector.subject = subject
		return expector
	}
}

type ExpectorFunc func(subject interface{}) *Expector

type Expector struct {
	t       Fataler
	subject interface{}
}

// ToBe checks simple equality, i.e. (x == y)
func (e Expector) ToBe(expected interface{}) {
	if e.subject != expected {
		e.t.Fatalf("%s is not %s", format(e.subject), format(expected))
	}
}

// ToEqual checks strict equality, i.e. (cmp.Diff(x, y) == "")
func (e Expector) ToEqual(expected interface{}) {
	diff := cmp.Diff(expected, e.subject)

	if diff != "" {
		e.t.Fatalf(diff)
	}
}

// NotToBe checks simple inequality, i.e. (x != y)
func (e Expector) NotToBe(expected interface{}) {
	if e.subject == expected {
		e.t.Fatalf("%s is %s", format(e.subject), format(expected))
	}
}

// NotToEqual checks strict inequality, i.e. (!cmp.Equal(x, y))
func (e Expector) NotToEqual(expected interface{}) {
	if cmp.Equal(e.subject, expected) {
		e.t.Fatalf("%s equals %s", format(e.subject), format(expected))
	}
}

func format(v interface{}) string {
	str, ok := v.(string)
	if ok {
		return fmt.Sprintf("'%s'", str)
	}
	return fmt.Sprint(v)
}
