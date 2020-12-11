package expectate_test

import (
	"testing"

	"github.com/gomagedon/expectate"
)

var zeroTests = []ExpectTest{
	{
		name:            "int zero is zero",
		subject:         int(0),
		expectedFailure: "",
	},
	{
		name:            "uint zero is zero",
		subject:         uint(0),
		expectedFailure: "",
	},
	{
		name:            "int8 zero is zero",
		subject:         int8(0),
		expectedFailure: "",
	},
	{
		name:            "uint8 zero is zero",
		subject:         uint8(0),
		expectedFailure: "",
	},
	{
		name:            "int16 zero is zero",
		subject:         int16(0),
		expectedFailure: "",
	},
	{
		name:            "uint16 zero is zero",
		subject:         uint16(0),
		expectedFailure: "",
	},
	{
		name:            "int32 zero is zero",
		subject:         int32(0),
		expectedFailure: "",
	},
	{
		name:            "uint32 zero is zero",
		subject:         uint32(0),
		expectedFailure: "",
	},
	{
		name:            "int64 zero is zero",
		subject:         int64(0),
		expectedFailure: "",
	},
	{
		name:            "uint64 zero is zero",
		subject:         uint64(0),
		expectedFailure: "",
	},
	{
		name:            "float32 zero is zero",
		subject:         float32(0),
		expectedFailure: "",
	},
	{
		name:            "float64 zero is zero",
		subject:         float64(0),
		expectedFailure: "",
	},
	{
		name:            "float32 0.0 is zero",
		subject:         float32(0.0),
		expectedFailure: "",
	},
	{
		name:            "float64 0.0 is zero",
		subject:         float64(0.0),
		expectedFailure: "",
	},
	{
		name:            "int 1 is not zero",
		subject:         int(1),
		expectedFailure: "1 is not zero\n",
	},
	{
		name:            "uint 1 is not zero",
		subject:         uint(1),
		expectedFailure: "1 is not zero\n",
	},
	{
		name:            "int8 1 is not zero",
		subject:         int8(1),
		expectedFailure: "1 is not zero\n",
	},
	{
		name:            "uint8 1 is not zero",
		subject:         uint8(1),
		expectedFailure: "1 is not zero\n",
	},
	{
		name:            "int16 1 is not zero",
		subject:         int16(1),
		expectedFailure: "1 is not zero\n",
	},
	{
		name:            "uint16 1 is not zero",
		subject:         uint16(1),
		expectedFailure: "1 is not zero\n",
	},
	{
		name:            "int32 1 is not zero",
		subject:         int32(1),
		expectedFailure: "1 is not zero\n",
	},
	{
		name:            "uint32 1 is not zero",
		subject:         uint32(1),
		expectedFailure: "1 is not zero\n",
	},
	{
		name:            "int64 1 is not zero",
		subject:         int64(1),
		expectedFailure: "1 is not zero\n",
	},
	{
		name:            "uint64 1 is not zero",
		subject:         uint64(1),
		expectedFailure: "1 is not zero\n",
	},
	{
		name:            "float32 1 is not zero",
		subject:         float32(1),
		expectedFailure: "1 is not zero\n",
	},
	{
		name:            "float64 1 is not zero",
		subject:         float64(1),
		expectedFailure: "1 is not zero\n",
	},
	{
		name:            "float32 with decimal is not zero",
		subject:         float32(1.5),
		expectedFailure: "1.5 is not zero\n",
	},
	{
		name:            "float64 with decimal is not zero",
		subject:         float64(1.5),
		expectedFailure: "1.5 is not zero\n",
	},
	{
		name:            "float32 0.1 is not zero",
		subject:         float32(0.1),
		expectedFailure: "0.1 is not zero\n",
	},
	{
		name:            "float64 0.1 is not zero",
		subject:         float64(0.1),
		expectedFailure: "0.1 is not zero\n",
	},
}

func TestZero(t *testing.T) {
	for _, test := range zeroTests {
		t.Run(test.name, func(t *testing.T) {
			mockTestingT := new(MockTestingT)
			expect := expectate.Expect(mockTestingT)

			expect(test.subject).ToBeZero()

			if mockTestingT.FataledWith != test.expectedFailure {
				t.Fatal("Expected:", test.expectedFailure,
					"\nGot:", mockTestingT.FataledWith)
			}
		})
	}
}
