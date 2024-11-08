package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {

	// create router
	mux := chi.NewRouter()

	// setup middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)

	// define application routes
	mux.Get("/", app.HomePage)
	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Post("/register", app.PostRegisterPage)
	mux.Get("/activate", app.ActivateAccount)

	// route with authentication required
	mux.Mount("/members", app.authRoute())

	// test route
	mux.Get("/test-email", func(w http.ResponseWriter, r *http.Request) {

		m := Mail{
			Domain:      "localhost",
			Host:        "localhost",
			Port:        1025,
			Encryption:  "none",
			FromAddress: "pj@pj.com",
			FromName:    "PJ",
			ErrorChan:   make(chan error),
		}

		msg := Message{
			To:      "testmail@test.com",
			Subject: "Test Mail",
			Data:    "Testing Email service",
		}

		m.sendMail(msg, make(chan error))
	})
	return mux
}

func (app *Config) authRoute() http.Handler {
	// create router
	mux := chi.NewRouter()

	// setup middleware
	mux.Use(app.Auth)

	mux.Get("/plans", app.ChooseSubscription)
	mux.Get("/subscribe", app.SubscribeToPlan)

	return mux
}
