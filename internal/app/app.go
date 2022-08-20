package app

import (
	"errors"
	"rrashidov/usesthisreader/internal/scheduler"
	"rrashidov/usesthisreader/internal/usesthisreader"
)

var (
	SchedulerNotProvidedError      = errors.New("Scheduler Not Provided!")
	UsesThisReaderNotProvidedError = errors.New("usesthis.com reader not provided")
)

type Application struct {
	s scheduler.Scheduler
	r usesthisreader.UsesThisReader
}

func NewApplication(s scheduler.Scheduler, r usesthisreader.UsesThisReader) *Application {
	return &Application{
		s: s,
		r: r,
	}
}

func (app Application) Run() error {
	if app.s == nil {
		return SchedulerNotProvidedError
	}

	if app.r == nil {
		return UsesThisReaderNotProvidedError
	}

	return app.s.Schedule(app.r)
}
