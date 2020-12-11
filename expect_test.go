package expectate_test

import "fmt"

type MockTestingT struct {
	FataledWith string
}

func (t *MockTestingT) Fatal(args ...interface{}) {
	t.FataledWith = fmt.Sprintln(args...)
}
