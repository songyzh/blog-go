package main

import (
	"flag"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/songyzh/blog/handlers"
	"github.com/songyzh/blog/middlewares"
	"net/http"
)

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api", func(r chi.Router){
		r.Route("/posts", func(r chi.Router){
			r.With(middlewares.Paginate).Get("/", handlers.MGetPosts)
			r.Get("/{slug:[a-z0-9-]+}", handlers.GetPostBySlug)
		})
		r.Route("/clean_cache", func(r chi.Router) {
			r.Get("/", handlers.CleanCache)
		})
		r.Route("/tags", func(r chi.Router) {
			r.Get("/", handlers.GetAvailableTags)
		})
	})

	r.Route("/benchmark", func(r chi.Router) {
		r.Get("/", handlers.Benchmark)
	})

	http.ListenAndServe(":3333", r)
}
