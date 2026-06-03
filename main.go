package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/iv-tunate/fiids/config"
	"github.com/iv-tunate/fiids/database"
	router "github.com/iv-tunate/fiids/routers"
	"github.com/iv-tunate/fiids/services"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer dbConn.Close()

	dbQueries := database.New(dbConn)

	apiCfg := config.ApiConfig{
		DB: dbQueries,
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	
	go services.ScrapeFeeds(ctx, dbQueries, 10, time.Minute) 
	router := router.SetupRouter(port, &apiCfg)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server is running on port %s", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to listen and serve with error/n%v", err)
	}
}
