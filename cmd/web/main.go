package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	port := flag.Int("port", 4000, "HTTP Port number")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		logger: logger,
	}

	logger.Info("starting server", "port", *port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), app.routes()) // web server
	logger.Error(err.Error())
	os.Exit(1)
}
