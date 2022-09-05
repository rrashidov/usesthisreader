package scheduler

import (
	"testing"
	"time"
)

func TestPeriodicScheduler(t *testing.T) {

	tests := []struct {
		name                       string
		pass_exec                  bool
		period                     int
		wait_before_check          int
		expectedEror               error
		expectedNumberOfExecutions int
	}{
		{"Null exec", false, 0, 0, ErrProvidedExecIsNil, 0},
		{"Non-positive period", true, 0, 0, ErrPeriodShouldBePositive, 0},
		{"Big period", true, 10000, 0, nil, 0},
		{"Normal period one execution", true, 10, 15, nil, 1},
		{"Smaller period two executions", true, 10, 25, nil, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PeriodicScheduler{
				period: tt.period,
			}

			exec := &TestLogic{}

			var err error

			go func() {
				if tt.pass_exec {
					err = s.Schedule(exec)
				} else {
					err = s.Schedule(nil)
				}
			}()

			time.Sleep(3 * time.Second)

			if err != tt.expectedEror {
				t.Errorf("Scheduler did not return proper error: expected %q, got: %q", tt.expectedEror, err)
			}

			time.Sleep(time.Duration(tt.wait_before_check) * time.Millisecond)

			if exec.execCount < tt.expectedNumberOfExecutions {
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
