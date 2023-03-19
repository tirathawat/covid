package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// App is the configuration for the application.
type App struct {
	Host     string `envconfig:"HOST" default:""`
	Port     string `envconfig:"PORT" default:"8080"`
	CovidURL string `envconfig:"COVID_URL" default:"https://static.wongnai.com/devinterview/covid-cases.json"`
}

// New returns a new configuration.
func New() *App {
	godotenv.Load()
	cfg := new(App)
	envconfig.MustProcess("", cfg)
	return cfg
}
