package main

import (
	"fmt"
	"log"
	"net/http"

	"cv-generator-services/internal/handler/cv_generator_handler"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	// Load the configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Create the router
	r := mux.NewRouter()

	// Create the handler and register the routes
	h := cv_generator_handler.NewHandler()
	h.RegisterRoutes(r)

	// Start the server
	addr := fmt.Sprintf(":%s", config.Port)
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func loadConfig() (config, error) {
	// Load the configuration
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read configuration: %v", err)
	}

	// Unmarshal the configuration into a struct
	var cfg config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %v", err)
	}

	return &cfg, nil
}
