package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/ShankarBirTamang/goProject/internal/app"
	"github.com/ShankarBirTamang/goProject/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	routes := routes.SetupRoutes(app)
	http.Handle("/", routes)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Application started on port %d", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
