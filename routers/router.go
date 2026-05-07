package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/iv-tunate/fiids/config"
	"github.com/iv-tunate/fiids/handlers"
	"github.com/iv-tunate/fiids/middleware"
)

func SetupRouter(port string, apiCfg *config.ApiConfig) *chi.Mux{
	router := chi.NewRouter()

	router.Use(middleware.LoggingMiddleware)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: false,
		MaxAge: 300,
		ExposedHeaders:    []string{"Link"},
		AllowedHeaders:   []string{"*"},
	}))

	v1Router := chi.NewRouter()	
	v1Router.Get("/healthz", handlers.HandlerHealth)
	v1Router.Get("/error", handlers.HandlerError)
	router.Mount("/v1", v1Router)

	cfg := handlers.New(apiCfg)

	v1Router.Post("/users", cfg.RegisterUser)

	return  router
}