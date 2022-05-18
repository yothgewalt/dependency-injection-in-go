package libraries

import (
	"log"

	"github.com/spf13/viper"
)

type Environment struct {
	EchoServerPort  string `mapstructure:"ECHO_SERVER_PORT"`
	EnvironmentMode string `mapstructure:"ENVIRONMENT_MODE"`
	LogOutput       string `mapstructure:"LOG_OUTPUT"`
}

func NewEnvironment() Environment {
	environment := Environment{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("cannot read configuration")
	}

	if err := viper.Unmarshal(&environment); err != nil {
		log.Fatalf("environment cannot be loaded: %v\n", err)
	}

	return environment
}
