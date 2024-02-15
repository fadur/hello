package cmd

import (
	"net/http"

	"github.com/fadur/hello/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func Server() error {

	r := chi.NewRouter()

	r.Get("/", handlers.IndexHandler)
	r.Get("/users", handlers.ListUsers)
	r.Post("/users", handlers.CreateUser)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return err
	}
	return nil

}
