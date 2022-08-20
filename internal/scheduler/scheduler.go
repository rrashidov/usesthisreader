package scheduler

import (
	"errors"
	"rrashidov/usesthisreader/internal/generic"
	"time"
)

var (
	ErrProvidedExecIsNil      = errors.New("provided exec is nil")
	ErrPeriodShouldBePositive = errors.New("period should be positive")
)

type Scheduler interface {
	Schedule(exec generic.GenericLogic) error
}

type PeriodicScheduler struct {
	period int
}

func (pr PeriodicScheduler) Schedule(exec generic.GenericLogic) error {

	if exec == nil {
		return ErrProvidedExecIsNil
	}

	if pr.period <= 0 {
		return ErrPeriodShouldBePositive
	}

	go func() {
		for {
			time.Sleep(time.Duration(pr.period) * time.Millisecond)

			exec.Run()
		}
	}()

	return nil
}
