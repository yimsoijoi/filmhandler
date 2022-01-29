package config

import (
	"github.com/spf13/viper"
	"github.com/yimsoijoi/filmhandler/lib/postgres"
)

type Config struct {
	Postgres postgres.Config `mapstructure:"postgres"`
	Padding  int             `mapstructure:"padding"`
}

func Load() (*Config, error) {
	var conf Config

	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		// Handler error
	}
	if err := viper.Unmarshal(&conf); err != nil {
		// Handle error
	}

	return &conf, nil
}
