package main

import (
	"flag"
	"fmt"
	"rrashidov/usesthisreader/internal/app"
	"rrashidov/usesthisreader/internal/scheduler"
	"rrashidov/usesthisreader/internal/usesthisreader"
)

const (
	DefaultExecutionPeriod      int    = 60 * 60 * 1000
	URL                         string = ""
	LocalStorageDefaultFilePath string = ".usesthisreader"
	AWSRegion                   string = ""
	Recipient                   string = ""
	Sender                      string = ""
)

type app_cfg struct {
	exec_period int
}

func main() {
	var cfg app_cfg

	flag.IntVar(&cfg.exec_period, "execPeriod", DefaultExecutionPeriod, "Execution period")

	flag.Parse()

	print_application_config(cfg)

	app, err := init_application(cfg)

	if err != nil {
		fmt.Println("Error initializing the application: " + err.Error())
	}

	err = app.Run()

	if err != nil {
		fmt.Println("Error running the application: " + err.Error())
	}
}

func init_application(cfg app_cfg) (*app.Application, error) {
	sched := scheduler.NewPeriodicScheduler(cfg.exec_period)

	reader := usesthisreader.NewUsesThisReader("", LocalStorageDefaultFilePath, AWSRegion, Recipient, Sender)

	return app.NewApplication(sched, reader), nil
}

func print_application_config(cfg app_cfg) {
	fmt.Printf("Starting application with the following config:")
	fmt.Printf("execution period: %d", cfg.exec_period)
}
