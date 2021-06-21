package main

import (
	"net/http"

	"github.com/agniadvani/go-course/pkg/config"
	handler "github.com/agniadvani/go-course/pkg/handlers"
	"github.com/bmizerany/pat"
)

//Using Pat Router
func router(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))
	return mux
}
