package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/aliemreipek/bookings/pkg/config"
	"github.com/aliemreipek/bookings/pkg/handlers"
	"github.com/aliemreipek/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

// portNumber is the number which my application runs in web browser
const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager
// main is the main application function
func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create the template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println("Application starting on port", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
