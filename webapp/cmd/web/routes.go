package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(app.addIPToContext)
	mux.Use(app.Session.LoadAndSave)
	mux.Get("/", app.Home)
	mux.Post("/login", app.Login)

	// mux.Get("/user/profile", app.Profile)

	mux.Route("/user", func(mux chi.Router) {
		mux.Use(app.auth)
		mux.Get("/profile", app.Profile)
	})

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
