package services

import (
	"objective-service/services/postgres"

	apiservices "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-api"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-app"
)

type Config struct {
	PgConfig  postgres.Config    `yaml:"postgres" mapstructure:"postgres" json:"postgres"`
	ApiConfig apiservices.Config `yaml:"api" mapstructure:"api" json:"api"`
}

func ProvideServices(cfg *Config) {
	core.Supply(&cfg.PgConfig, &cfg.ApiConfig)
	core.Provides(postgres.NewService, apiservices.NewService, apiservices.NewRouter, apiservices.NewMetricsReporter)

}
