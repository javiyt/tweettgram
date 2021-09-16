package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	BotToken     string `required:"true" split_words:"true"`
	ValidChannel int64  `required:"true" split_words:"true"`
}

func NewEnvConfig() EnvConfig {
	var e EnvConfig
	err := envconfig.Process("", &e)
	if err != nil {
		log.Fatal(err.Error())
	}

	return e
}

func (ec EnvConfig) IsValidChannel(channelID int64) bool {
	return ec.ValidChannel == channelID
}
