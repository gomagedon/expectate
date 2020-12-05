package expect_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/steve-kaufman/go-expect"
)

type MockTestingT struct {
	FataledWith string
}

func (t *MockTestingT) Fatal(args ...interface{}) {
	t.FataledWith = fmt.Sprintln(args...)
}

type Person struct {
	Name     string
	Age      int
	Job      string
	Birthday time.Time
}

var myPointerToPerson = &Person{
	Name:     "John Doe",
	Age:      30,
	Job:      "Electrician",
	Birthday: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
}

type ToBeTest struct {
	name            string
	subject         interface{}
	object          interface{}
	expectedFailure string
}

var toBeTests = []ToBeTest{
	{
		name:            "2 is 2",
		subject:         2,
		object:          2,
		expectedFailure: "",
	},
	{
		name:            "2 is not 3",
		subject:         2,
		object:          3,
		expectedFailure: "2 is not 3\n",
	},
	{
		name:            "'foo' is 'foo'",
		subject:         "foo",
		object:          "foo",
		expectedFailure: "",
	},
	{
		name:            "'foo' is not 'bar'",
		subject:         "foo",
		object:          "bar",
		expectedFailure: "foo is not bar\n",
	},
	{
		name:            "0 is 0",
		subject:         0,
		object:          0,
		expectedFailure: "",
	},
	{
		name:            "0 is not ''",
		subject:         0,
		object:          "",
		expectedFailure: "0 is not \n",
	},
	{
		name:            "0 is not nil",
		subject:         0,
		object:          nil,
		expectedFailure: "0 is not <nil>\n",
	},
	{
		name:            "pointer to struct is itself",
		subject:         myPointerToPerson,
		object:          myPointerToPerson,
		expectedFailure: "",
	},
	{
		name:    "pointer to struct is not copy of struct",
		subject: myPointerToPerson,
		object: Person{
			Name:     myPointerToPerson.Name,
			Age:      myPointerToPerson.Age,
			Job:      myPointerToPerson.Job,
			Birthday: myPointerToPerson.Birthday,
		},
		expectedFailure: "&{John Doe 30 Electrician 1990-01-01 00:00:00 +0000 UTC} is not {John Doe 30 Electrician 1990-01-01 00:00:00 +0000 UTC}\n",
	},
	{
		name:    "pointer to struct is not pointer to copy of struct",
		subject: myPointerToPerson,
		object: &Person{
			Name:     myPointerToPerson.Name,
			Age:      myPointerToPerson.Age,
			Job:      myPointerToPerson.Job,
			Birthday: myPointerToPerson.Birthday,
		},
		expectedFailure: "&{John Doe 30 Electrician 1990-01-01 00:00:00 +0000 UTC} is not &{John Doe 30 Electrician 1990-01-01 00:00:00 +0000 UTC}\n",
	},
	{
		name: "struct is copy of struct",
		subject: Person{
			Name:     myPointerToPerson.Name,
			Age:      myPointerToPerson.Age,
			Job:      myPointerToPerson.Job,
			Birthday: myPointerToPerson.Birthday,
		},
		object: Person{
			Name:     myPointerToPerson.Name,
			Age:      myPointerToPerson.Age,
			Job:      myPointerToPerson.Job,
			Birthday: myPointerToPerson.Birthday,
		},
		expectedFailure: "",
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
		expectedFailure: "{John Doe 30 Electrician 2000-01-01 00:00:00 +0000 UTC} is not {John Smith 30 Electrician 2000-01-01 00:00:00 +0000 UTC}\n",
	},
}

func TestToBe(t *testing.T) {
	for _, test := range toBeTests {
		t.Run(test.name, func(t *testing.T) {
			mockTestingT := new(MockTestingT)
			expect := expect.Expect(mockTestingT)

			expect(test.subject).ToBe(test.object)

			if mockTestingT.FataledWith != test.expectedFailure {
				t.Fatal("Expected:", test.expectedFailure,
					"\nGot:", mockTestingT.FataledWith)
			}
		})
	}
}
