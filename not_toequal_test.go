package expectate_test

import (
	"testing"
	"time"

	"github.com/gomagedon/expectate"
)

var notToEqualTests = []ExpectTest{
	{
		name:            "2 equals 2",
		subject:         2,
		object:          2,
		expectedFailure: "2 equals 2\n",
	},
	{
		name:            "2 does not equal 3",
		subject:         2,
		object:          3,
		expectedFailure: "",
	},
	{
		name:            "'foo' equals 'foo'",
		subject:         "foo",
		object:          "foo",
		expectedFailure: "foo equals foo\n",
	},
	{
		name:            "'foo' does not equal 'bar'",
		subject:         "foo",
		object:          "bar",
		expectedFailure: "",
	},
	{
		name: "pointer to struct equals pointer to copy of struct",
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
		expectedFailure: "&{John Doe 30 Electrician 1990-01-01 00:00:00 +0000 UTC} equals &{John Doe 30 Electrician 1990-01-01 00:00:00 +0000 UTC}\n",
	},
}

func TestNotToEqual(t *testing.T) {
	for _, test := range notToEqualTests {
		t.Run(test.name, func(t *testing.T) {
			mockTestingT := new(MockTestingT)
			expect := expectate.Expect(mockTestingT)

			expect(test.subject).NotToEqual(test.object)

			if mockTestingT.FataledWith != test.expectedFailure {
				t.Fatal("Expected:", test.expectedFailure,
					"\nGot:", mockTestingT.FataledWith)
			}
		})
	}
}
