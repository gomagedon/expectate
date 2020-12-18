package expectate_test

import (
	"testing"

	"github.com/gomagedon/expectate"
)

var notEmptyStrTests = []ExpectTest{
	{
		name:            "empty string is empty string",
		subject:         "",
		expectedFailure: "this is an empty string\n",
	},
	{
		name:            "'foo' is not empty string",
		subject:         "foo",
		expectedFailure: "",
	},
	{
		name:            "zero is not empty string",
		subject:         0,
		expectedFailure: "",
	},
	{
		name:            "3 is not empty string",
		subject:         3,
		expectedFailure: "",
	},
	{
		name:            "a struct is not empty string",
		subject:         *samplePointerToPerson,
		expectedFailure: "",
	},
	{
		name:            "a pointer to a struct is not empty string",
		subject:         samplePointerToPerson,
		expectedFailure: "",
	},
	{
		name:            "an empty struct is not empty string",
		subject:         Person{},
		expectedFailure: "",
	},
}

func TestNotEmptyStr(t *testing.T) {
	for _, test := range notEmptyStrTests {
		t.Run(test.name, func(t *testing.T) {
			mockTestingT := new(MockTestingT)
			expect := expectate.Expect(mockTestingT)

			expect(test.subject).NotToBeEmptyStr()

			if mockTestingT.FataledWith != test.expectedFailure {
				t.Fatal("Expected:", test.expectedFailure,
					"\nGot:", mockTestingT.FataledWith)
			}
		})
	}
}
