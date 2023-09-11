package config

import (
    "fmt"
    "os"

)

type Config struct {
    Port     string
    Database struct {
        Host     string
        Port     string
        Username string
        Password string
        Name     string
    }
}
