package app

import (
	"errors"
	"rrashidov/usesthisreader/internal/generic"
	"rrashidov/usesthisreader/internal/scheduler"
	"rrashidov/usesthisreader/internal/usesthisreader"
	"testing"
)

var (
	ErrorForTestingPurposes = errors.New("test Scheduler Error")
)

func TestApp(t *testing.T) {

	r := &TestReader{}

	s := &TestScheduler{}

	errorSched := &TestScheduler{
		err: ErrorForTestingPurposes,
	}

	tests := []struct {
		name              string
		s                 scheduler.Scheduler
		r                 usesthisreader.UsesThisReader
		expectedError     error
		expectedSchedules int
	}{
		{"Scheduler not provided", nil, r, SchedulerNotProvidedError, 0},
		{"usesthis.com reader not provided", s, nil, UsesThisReaderNotProvidedError, 0},
		{"Scheduler returns error", errorSched, r, ErrorForTestingPurposes, 0},
		{"Happy path", s, r, nil, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Application{
				s: tt.s,
				r: tt.r,
			}

			err := a.Run()

			if err != tt.expectedError {
				t.Errorf("Application run expected to return %s, but returned %s", tt.expectedError, err)
			}

			if s.schedules != tt.expectedSchedules {
				t.Errorf("Application should have been scheduled; expected schedules: %d, got: %d", tt.expectedSchedules, s.schedules)
			}
		})
	}
}

type TestScheduler struct {
	schedules int
	err       error
}

func (ts *TestScheduler) Schedule(exec generic.GenericLogic) error {
	ts.schedules++

	return ts.err
}

type TestReader struct{}

func (tr TestReader) Run() error {
	return nil
}
