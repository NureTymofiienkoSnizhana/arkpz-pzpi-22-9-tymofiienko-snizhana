package api

import (
	"context"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/api/handlers"
	"github.com/NureTymofiienkoSnizhana/arkpz-pzpi-22-9-tymofiienko-snizhana/Pract1/arkpz-pzpi-22-9-tymofiienko-snizhana-task2/src/data"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type Config struct {
	MasterDB data.MasterDB
}

func Run(config Config) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), handlers.MasterDBContextKey, config.MasterDB)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/petandhealth", func(r chi.Router) {
		r.Route("/v1/public", func(r chi.Router) {
			r.Get("/auth", handlers.Auth)
			r.Post("/registration", handlers.Registration)
			r.Post("/adddevice", handlers.AddDevice)
			r.Post("/addpet", handlers.AddPet)
			r.Put("/updateuser", handlers.UpdateUser)
			r.Put("/updatepet", handlers.UpdatePet)
			r.Put("/updatedevice", handlers.UpdateDevice)
		})
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
