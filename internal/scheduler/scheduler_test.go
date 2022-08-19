package scheduler

import (
	"rrashidov/usesthisreader/internal/generic"
	"testing"
	"time"
)

func TestPeriodicScheduler(t *testing.T) {

	exec := &TestLogic{}

	tests := []struct {
		name                       string
		exec                       generic.GenericLogic
		period                     int
		wait_before_check          int
		expectedEror               error
		expectedNumberOfExecutions int
	}{
		{"Null exec", nil, 0, 0, ErrProvidedExecIsNil, 0},
		{"Non-positive period", exec, 0, 0, ErrPeriodShouldBePositive, 0},
		{"Big period", exec, 10000, 0, nil, 0},
		{"Normal period, one execution", exec, 10, 15, nil, 1},
		{"Smaller period, two executions", exec, 10, 25, nil, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PeriodicScheduler{
				period: tt.period,
			}

			err := s.Schedule(tt.exec)

			if err != tt.expectedEror {
				t.Errorf("Scheduler did not return proper error: expected %q, got: %q", tt.expectedEror, err)
			}

			time.Sleep(time.Duration(tt.wait_before_check) * time.Millisecond)

			if tt.expectedNumberOfExecutions < exec.execCount {
				t.Errorf("Scheduler did not execute provided logic; expected: %d, got: %d", tt.expectedNumberOfExecutions, exec.execCount)
			}
		})

	}
}

type TestLogic struct {
	execCount int
	err       error
}

func (tl *TestLogic) Run() error {
	tl.execCount++

	return tl.err
}
