package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/mambo-dev/kenya-locations/config"
	"github.com/mambo-dev/kenya-locations/internal/database"
	handler "github.com/mambo-dev/kenya-locations/internal/handlers"
	"github.com/mambo-dev/kenya-locations/internal/middleware"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.Connect(cfg.DBURL, cfg.TAuthToken)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	router.Use(middleware.SecureHeaders)

	v1 := chi.NewRouter()

	router.Mount("/v1", v1)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 30 * time.Second,
	}

	serverHandlers := handler.NewHandler(db, cfg)

	v1.Get("/counties", serverHandlers.GetCounties)
	v1.Get("/counties/{countyID}", serverHandlers.GetCounty)
	v1.Get("/counties/{countyID}/sub-counties", serverHandlers.GetCountySubCounties)

	v1.Get("/sub-counties", serverHandlers.GetSubCounties)
	v1.Get("/sub-counties/{subCountyID}", serverHandlers.GetSubCounty)
	v1.Get("/sub-counties/{subCountyID}/wards", serverHandlers.GetSubCountyWards)

	v1.Get("/wards", serverHandlers.GetWards)
	v1.Get("/wards/{wardID}", serverHandlers.GetWard)

	log.Printf("Server running on http://localhost:%s", cfg.Port)
	log.Fatal(srv.ListenAndServe())
}
