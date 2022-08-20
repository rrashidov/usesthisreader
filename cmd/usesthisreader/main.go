package main

import (
	"fmt"
	"rrashidov/usesthisreader/internal/app"
	"rrashidov/usesthisreader/internal/scheduler"
	"rrashidov/usesthisreader/internal/usesthisreader"
)

const (
	ExecutionPeriod int    = 60 * 60 * 1000
	URL             string = ""
)

func main() {
	app, err := init_application()

	if err != nil {
		fmt.Println("Error initializing the application: " + err.Error())
	}

	err = app.Run()

	if err != nil {
		fmt.Println("Error running the application: " + err.Error())
	}
}

func init_application() (*app.Application, error) {
	sched := scheduler.NewPeriodicScheduler(ExecutionPeriod)

	// TODO: pass uses this API url
	reader := usesthisreader.NewUsesThisReader("", ".usesthisreader")

	return app.NewApplication(sched, reader), nil
}
