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

	r.Route("/api/pet-and-health", func(r chi.Router) {
		r.Route("/v1/login", func(r chi.Router) {
			r.Get("/auth", handlers.Auth)
			r.Post("/registration", handlers.Registration)
		})
		r.Route("/v1/admin/pets", func(r chi.Router) {
			r.Post("/add-pet", handlers.AddPet)
			r.Delete("/delete-pet", handlers.DeletePet)
			r.Put("/update-pet", handlers.UpdatePet)
		})
		r.Route("/v1/admin/devices", func(r chi.Router) {
			r.Post("/add-device", handlers.AddDevice)
			r.Put("/update-device", handlers.UpdateDevice)
		})
		r.Route("/v1/user", func(r chi.Router) {
			r.Put("/update-user", handlers.UpdateUser)
			r.Get("/pets", handlers.GetPets)
			r.Get("/pet-info", handlers.PetInfo)
		})
		r.Route("/v1/vet", func(r chi.Router) {
			//r.Get("/pet-report", handlers.PetReport)
		})
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
