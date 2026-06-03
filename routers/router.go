package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/iv-tunate/fiids/config"
	"github.com/iv-tunate/fiids/handlers"
	"github.com/iv-tunate/fiids/middleware"
)

func SetupRouter(port string, apiCfg *config.ApiConfig) *chi.Mux{
	cfg := handlers.New(apiCfg)
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
	
	router.Mount("/v1", v1Router)
	v1Router.Use(middleware.ApiKeyAuthMiddleware(apiCfg))


	router.Get("/healthz", handlers.HandlerHealth)
	router.Get("/error", handlers.HandlerError)
	
	router.Post("/users", cfg.RegisterUser)
	router.Get("/users", cfg.GetUsers)
	router.Get("/user", cfg.GetUserById)
	router.Post("/auth/apikey", cfg.GenerateApiKey)
	
	v1Router.Post("/feed", cfg.CreateFeed)
	v1Router.Get("/feeds", cfg.GetFeeds)
	v1Router.Post("/follow_feed", cfg.FollowFeed)
	v1Router.Get("/followed_feeds", cfg.GetFollowedFeeds)
	v1Router.Delete("/unfollow_feed/{id}", cfg.UnfollowFeed)
	return  router
}