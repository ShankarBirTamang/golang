package app

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

// NewApplication creates a new application instance
func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		Logger: logger,
	}

	return app, nil
}
