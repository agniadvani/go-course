package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/agniadvani/go-course/pkg/config"
	handler "github.com/agniadvani/go-course/pkg/handlers"
	render "github.com/agniadvani/go-course/pkg/renders"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {
	app.InProduction = false

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

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
