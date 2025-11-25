package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ShankarBirTamang/goProject/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

// NewApplication creates a new application instance
func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		Logger:         logger,
		WorkoutHandler: api.NewWorkoutHandler(),
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
