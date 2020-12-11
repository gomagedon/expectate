package expectate_test

import (
	"testing"

	"github.com/gomagedon/expectate"
)

var emptyStrTests = []ExpectTest{
	{
		name:            "empty string is empty string",
		subject:         "",
		expectedFailure: "",
	},
	{
		name:            "'foo' is not empty string",
		subject:         "foo",
		expectedFailure: "foo is not an empty string\n",
	},
	{
		name:            "zero is not empty string",
		subject:         0,
		expectedFailure: "0 is not an empty string\n",
	},
	{
		name:            "3 is not empty string",
		subject:         3,
		expectedFailure: "3 is not an empty string\n",
	},
	{
		name:            "a struct is not empty string",
		subject:         *samplePointerToPerson,
		expectedFailure: "{John Doe 30 Electrician 1990-01-01 00:00:00 +0000 UTC} is not an empty string\n",
	},
	{
		name:            "a pointer to a struct is not empty string",
		subject:         samplePointerToPerson,
		expectedFailure: "&{John Doe 30 Electrician 1990-01-01 00:00:00 +0000 UTC} is not an empty string\n",
	},
	{
		name:            "an empty struct is not empty string",
		subject:         Person{},
		expectedFailure: "{ 0  0001-01-01 00:00:00 +0000 UTC} is not an empty string\n",
	},
}

func TestEmptyStr(t *testing.T) {
	for _, test := range emptyStrTests {
		t.Run(test.name, func(t *testing.T) {
			mockTestingT := new(MockTestingT)
			expect := expectate.Expect(mockTestingT)

			expect(test.subject).ToBeEmptyStr()

			if mockTestingT.FataledWith != test.expectedFailure {
				t.Fatal("Expected:", test.expectedFailure,
					"\nGot:", mockTestingT.FataledWith)
			}
		})
	}
}
