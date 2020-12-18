package expectate_test

import (
	"testing"
	"time"

	"github.com/gomagedon/expectate"
	"github.com/google/go-cmp/cmp"
)

var toEqualTests = []ExpectTest{
	{
		name:            "2 equals 2",
		subject:         2,
		object:          2,
		expectedFailure: "",
	},
	{
		name:            "2 does not equal 3",
		subject:         2,
		object:          3,
		expectedFailure: cmp.Diff(3, 2) + "\n",
	},
	{
		name:            "'foo' equals 'foo'",
		subject:         "foo",
		object:          "foo",
		expectedFailure: "",
	},
	{
		name:            "'foo' does not equal 'bar'",
		subject:         "foo",
		object:          "bar",
		expectedFailure: cmp.Diff("bar", "foo") + "\n",
	},
	{
		name: "pointer to struct is pointer to copy of struct",
		subject: &Person{
			Name:     "John Doe",
			Age:      30,
			Job:      "Electrician",
			Birthday: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		object: &Person{
			Name:     "John Doe",
			Age:      30,
			Job:      "Electrician",
			Birthday: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		expectedFailure: "",
	},
}

func TestToEqual(t *testing.T) {
	for _, test := range toEqualTests {
		t.Run(test.name, func(t *testing.T) {
			mockTestingT := new(MockTestingT)
			expect := expectate.Expect(mockTestingT)

			expect(test.subject).ToEqual(test.object)

			if mockTestingT.FataledWith != test.expectedFailure {
				t.Fatal("Expected:", test.expectedFailure,
					"\nGot:", mockTestingT.FataledWith)
			}
		})
	}
}
