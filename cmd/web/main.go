package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agniadvani/go-course/pkg/config"
	handler "github.com/agniadvani/go-course/pkg/handlers"
	render "github.com/agniadvani/go-course/pkg/renders"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	r := handler.NewRepo(&app)
	handler.NewHandler(r)
	render.NewTemplate(&app)

	app.UseCache = false

	fmt.Println("Application running on port", portNumber)

	//Serving at port :8080 and routing using pat router
	srv := &http.Server{
		Addr:    portNumber,
		Handler: router(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
