package config

import (
	"log"
	"strings"
	"sync"

	"github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
)

var (
	serverConfig ServerConfig
	configOnce   sync.Once
)

type configBuilder struct {
	dotenvFile           string
	errorOnMissingDotenv bool
}

func ConfigBuilder() configBuilder {
	return configBuilder{}
}

func (b configBuilder) WithDotenvFile(file string) configBuilder {
	b.dotenvFile = file
	return b
}

func (b configBuilder) PanicOnMissingDotenv(status bool) configBuilder {
	b.errorOnMissingDotenv = status
	return b
}

func (b configBuilder) Build() ServerConfig {
	serverConfig = NewConfig()

	dotenvFile := ".env"
	if b.dotenvFile != "" {
		dotenvFile = b.dotenvFile
	}
	dotenvFeeder := feeder.DotEnv{Path: dotenvFile}
	envFeeder := feeder.Env{}

	err := config.New().AddStruct(&serverConfig).AddFeeder(dotenvFeeder).Feed()
	if err != nil {
		if strings.Contains(err.Error(), "no such file") && b.errorOnMissingDotenv {
			log.Fatalf("error loading config from dotenv file %s: %s", dotenvFile, err.Error())
		}
	}
	err = config.New().AddStruct(&serverConfig).AddFeeder(envFeeder).Feed()
	if err != nil {
		log.Fatalf("error loading config from environment: %s", err.Error())
	}

	serverConfig.postConfigSetup()

	return serverConfig
}

// Setup PostConfigEnv which have logic in them eg. SteamAPIEnabled = ?
func (s *ServerConfig) postConfigSetup() {
	// Check if there is an API key given --> set bool
	s.GridAPIEnabled = s.SteamGridDBAPIKey != ""

	// TODO: Make the datapath check in here
}

func Config() ServerConfig {
	configOnce.Do(func() {
		serverConfig = ConfigBuilder().Build()
	})

	return serverConfig
}

// TODO: Make datapath pre check so it doesn't have to end with "/"
type ServerConfig struct {
	ApiSecret     string `env:"API_SECRET"`
	TokenLifespan string `env:"JWT_TOKEN_LIFESPAN"` // Is given in hours

	DataPath string `env:"DATA_PATH"` // Has to end with a "/" for now

	SteamGridDBAPIKey string `env:"STEAMGRIDDB_API_KEY"`
	GridAPIEnabled    bool

	DebugEnabled bool `env:"DEBUG"`
}

// Bootstrap the applicatoin config struct with the default config
func NewConfig() ServerConfig {
	return ServerConfig{
		ApiSecret:         "yourapikey",
		TokenLifespan:     "10",
		DataPath:          "data/",
		SteamGridDBAPIKey: "",
		GridAPIEnabled:    false,
		DebugEnabled:      false,
	}
}
