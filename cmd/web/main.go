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

	mux := http.NewServeMux() // router - serveMux

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("POST /", app.sendTicketPost)

	logger.Info("starting server", "port", *port)

	// web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)
	logger.Error(err.Error())
	os.Exit(1)
}
