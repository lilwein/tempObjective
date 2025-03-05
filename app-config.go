package main

import (
	_ "embed"
	"objective-service/api"
	"objective-service/services"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-app"
	"github.com/rs/zerolog/log"
)

//go:embed config.yml
var projectConfigFile []byte

const ConfigFileEnvVar = "demo"

type Config struct {
	ServicesConfig services.Config `yaml:"services" mapstructure:"services"`
	ApiConfig      api.Config      `yaml:"api" mapstructure:"api"`
}

func init() {
	core.AppName = api.AppName
	var applicationConfig *Config
	if err := core.ReadConfig(string(projectConfigFile), ConfigFileEnvVar, &applicationConfig); err != nil {
		log.Fatal().Err(err).Msg("Failed to read config")
	}
	core.Supply(&applicationConfig.ApiConfig)
	services.ProvideServices(&applicationConfig.ServicesConfig)

}
