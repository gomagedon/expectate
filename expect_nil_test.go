package expectate_test

import (
	"fmt"
	"testing"

	"github.com/gomagedon/expectate"
)

type nilStringer struct {
	someField string
}

func (nilStringer) String() string {
	return "<nil>"
}

var nilTests = []ExpectTest{
	{
		name:            "nil is nil",
		subject:         nil,
		expectedFailure: "",
	},
	{
		name: "nil error is nil",
		subject: func() error {
			return nil
		}(),
		expectedFailure: "",
	},
	{
		name: "nil func is nil",
		subject: func() func() {
			return nil
		}(),
		expectedFailure: "",
	},
	{
		name:            "string '<nil>' is not nil",
		subject:         "<nil>",
		expectedFailure: "<nil> is not nil\n",
	},
	{
		name:            "nilStringer is not nil",
		subject:         nilStringer{},
		expectedFailure: "<nil> is not nil\n",
	},
	{
		name: "nil fmt.Stringer is nil",
		subject: func() fmt.Stringer {
			return nil
		}(),
		expectedFailure: "",
	},
}

func TestNil(t *testing.T) {
	for _, test := range nilTests {
		t.Run(test.name, func(t *testing.T) {
			mockTestingT := new(MockTestingT)
			expect := expectate.Expect(mockTestingT)

			expect(test.subject).ToBeNil()

			if mockTestingT.FataledWith != test.expectedFailure {
				t.Fatal("Expected:", test.expectedFailure,
					"\nGot:", mockTestingT.FataledWith)
			}
		})
	}
}
