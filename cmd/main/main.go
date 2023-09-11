package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/spf13/viper"

    "cv-generator-services/internal/handler"
)

func main() {
    // Load the configuration
    cfg, err := loadConfig()
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }

    // Create the router
    r := mux.NewRouter()

    // Create the handler and register the routes
    h := handler.NewHandler()
    h.RegisterRoutes(r)

    // Start the server
    addr := fmt.Sprintf(":%s", cfg.Port)
    log.Printf("Starting server on %s", addr)
    if err := http.ListenAndServe(addr, r); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}

func loadConfig() (*Config, error) {
    // Load the configuration
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    viper.SetConfigType("yaml")

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    // Unmarshal the configuration into a struct
    var cfg Config
    if err := viper.Unmarshal(&cfg); err != nil {
        return nil, err
    }

    return &cfg, nil
}