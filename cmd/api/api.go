package main

import (
	"net/http"
	"time"

	"github.com/apoorv-x12/social-media-app/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage

}

type config struct {
	address string
}

func (app *application) route() http.Handler{
		r := chi.NewRouter()

		// A good base middleware stack
		r.Use(middleware.RequestID)
		r.Use(middleware.RealIP)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)

		// Set a timeout value on the request context (ctx), that will signal
		// through ctx.Done() that the request has timed out and further
		// processing should be stopped.
		r.Use(middleware.Timeout(60 * time.Second))

		r.Get("/",app.GetHealth)
		
		return r
}

func (app *application) serve(router http.Handler) error {
	
	s := &http.Server{
		Addr:         app.config.address,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	return s.ListenAndServe()
}