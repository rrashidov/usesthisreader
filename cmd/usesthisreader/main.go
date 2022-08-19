package main

import (
	"fmt"
	"rrashidov/usesthisreader/internal/app"
)

func main() {
	app := &app.Application{}

	err := app.Run()

	if err != nil {
		fmt.Println("Error running the application: " + err.Error())
	}
}
