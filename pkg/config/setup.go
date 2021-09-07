package config

import (
	"github.com/hans-m-song/archive-ingest/pkg/config"
	"github.com/spf13/viper"
)

const (
	ServerHost = "SERVER_HOST"
	ServerPort = "SERVER_PORT"
)

func AugmentEnvSetup() {
	defaults := []config.EnvSpec{
		{Path: ServerPort, DefaultValue: "8000"},
		{Path: ServerHost, DefaultValue: "localhost"},
	}

	for _, spec := range defaults {
		viper.SetDefault(spec.Path, spec.DefaultValue)
	}

	config.Setup()
}

func Expose() {
	config.Expose()
}
