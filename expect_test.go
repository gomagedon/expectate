package expectate_test

import (
	"fmt"
	"time"
)

type MockTestingT struct {
	FataledWith string
}

func (t *MockTestingT) Fatalf(format string, args ...interface{}) {
	t.FataledWith = fmt.Sprintf(format, args...)
}

type Person struct {
	Name     string
	Age      int
	Job      string
	Birthday time.Time
}

var samplePointerToPerson = &Person{
	Name:     "John Doe",
	Age:      30,
	Job:      "Electrician",
	Birthday: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
}

type ExpectTest struct {
	name            string
	subject         interface{}
	object          interface{}
	expectedFailure string
}
