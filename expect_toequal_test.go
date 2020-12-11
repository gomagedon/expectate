package expectate_test

import (
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/google/go-cmp/cmp"
)

type ToEqualTest struct {
	name            string
	subject         interface{}
	object          interface{}
	expectedFailure string
}

var toEqualTests = []ToEqualTest{
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
			Name:     myPointerToPerson.Name,
			Age:      myPointerToPerson.Age,
			Job:      myPointerToPerson.Job,
			Birthday: myPointerToPerson.Birthday,
		},
		object: &Person{
			Name:     myPointerToPerson.Name,
			Age:      myPointerToPerson.Age,
			Job:      myPointerToPerson.Job,
			Birthday: myPointerToPerson.Birthday,
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
