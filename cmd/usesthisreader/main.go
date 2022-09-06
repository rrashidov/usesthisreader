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
	DefaultURL                  string = "https://usesthis.com/api/interviews/"
	DefaultLocalStorageFilePath string = ".usesthisreader"
	DefaultAWSRegion            string = "us-east-1"
)

type app_cfg struct {
	exec_period        int
	url                string
	local_storage_path string
	aws_region         string
	recipient          string
	sender             string
}

func main() {
	var cfg app_cfg

	flag.IntVar(&cfg.exec_period, "execPeriod", DefaultExecutionPeriod, "Execution period")
	flag.StringVar(&cfg.url, "url", DefaultURL, "API URL")
	flag.StringVar(&cfg.local_storage_path, "local-storage-path", DefaultLocalStorageFilePath, "Local storage path")
	flag.StringVar(&cfg.aws_region, "aws-region", DefaultAWSRegion, "AWS Region")
	flag.StringVar(&cfg.recipient, "recipient", "", "Mail notification recipient")
	flag.StringVar(&cfg.sender, "sender", "", "Mail notification sender")

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

	reader := usesthisreader.NewUsesThisReader(
		cfg.url,
		cfg.local_storage_path,
		cfg.aws_region,
		cfg.recipient,
		cfg.sender)

	return app.NewApplication(sched, reader), nil
}

func print_application_config(cfg app_cfg) {
	fmt.Printf("Starting application with the following config:\n")
	fmt.Printf("execution period: %d\n", cfg.exec_period)
	fmt.Printf("API url: %q\n", cfg.url)
	fmt.Printf("AWS region: %q\n", cfg.aws_region)
	fmt.Printf("Recipient: %q\n", cfg.recipient)
	fmt.Printf("Sender: %q\n", cfg.sender)
}
