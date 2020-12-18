package expectate_test

import (
	"testing"
	"time"

	"github.com/gomagedon/expectate"
)

var notToBeTests = []ExpectTest{
	{
		name:            "2 is 2",
		subject:         2,
		object:          2,
		expectedFailure: "2 is 2\n",
	},
	{
		name:            "2 is not 3",
		subject:         2,
		object:          3,
		expectedFailure: "",
	},
	{
		name:            "'foo' is 'foo'",
		subject:         "foo",
		object:          "foo",
		expectedFailure: "foo is foo\n",
	},
	{
		name:            "'foo' is not 'bar'",
		subject:         "foo",
		object:          "bar",
		expectedFailure: "",
	},
	{
		name:            "0 is 0",
		subject:         0,
		object:          0,
		expectedFailure: "0 is 0\n",
	},
	{
		name:            "0 is not ''",
		subject:         0,
		object:          "",
		expectedFailure: "",
	},
	{
		name:            "0 is not nil",
		subject:         0,
		object:          nil,
		expectedFailure: "",
	},
	{
		name:            "pointer to struct is itself",
		subject:         samplePointerToPerson,
		object:          samplePointerToPerson,
		expectedFailure: "&{John Doe 30 Electrician 1990-01-01 00:00:00 +0000 UTC} is &{John Doe 30 Electrician 1990-01-01 00:00:00 +0000 UTC}\n",
	},
	{
		name:            "pointer to struct is not copy of struct",
		subject:         samplePointerToPerson,
		object:          *samplePointerToPerson,
		expectedFailure: "",
	},
	{
		name: "pointer to struct is not pointer to copy of struct",
		subject: &Person{
			Name:     "Philip Fry",
			Age:      25,
			Job:      "Delivery Boy",
			Birthday: time.Date(1980, time.July, 7, 0, 0, 0, 0, time.UTC),
		},
		object: &Person{
			Name:     "Philip Fry",
			Age:      25,
			Job:      "Delivery Boy",
			Birthday: time.Date(1980, time.July, 7, 0, 0, 0, 0, time.UTC),
		},
		expectedFailure: "",
	},
	{
		name: "struct is copy of struct",
		subject: Person{
			Name:     "Hermes Conrad",
			Age:      38,
			Job:      "Beaurocrat",
			Birthday: time.Date(2967, time.August, 8, 0, 0, 0, 0, time.UTC),
		},
		object: Person{
			Name:     "Hermes Conrad",
			Age:      38,
			Job:      "Beaurocrat",
			Birthday: time.Date(2967, time.August, 8, 0, 0, 0, 0, time.UTC),
		},
		expectedFailure: "{Hermes Conrad 38 Beaurocrat 2967-08-08 00:00:00 +0000 UTC} is {Hermes Conrad 38 Beaurocrat 2967-08-08 00:00:00 +0000 UTC}\n",
	},
	{
		name: "struct is not struct with different values",
		subject: Person{
			Name:     "John Doe",
			Age:      30,
			Job:      "Electrician",
			Birthday: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		object: Person{
			Name:     "John Smith",
			Age:      30,
			Job:      "Electrician",
			Birthday: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		expectedFailure: "",
	},
}

func TestNotToBe(t *testing.T) {
	for _, test := range notToBeTests {
		t.Run(test.name, func(t *testing.T) {
			mockTestingT := new(MockTestingT)
			expect := expectate.Expect(mockTestingT)

			expect(test.subject).NotToBe(test.object)

			if mockTestingT.FataledWith != test.expectedFailure {
				t.Fatal("Expected:", test.expectedFailure,
					"\nGot:", mockTestingT.FataledWith)
			}
		})
	}
}
