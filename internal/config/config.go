package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string           `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string           `yaml:"storage_path" env-required:"true"`
	HttpServer  HttpServerConfig `yaml:"http_server"`
}

type HttpServerConfig struct {
	Address string `yaml:"address"`
}



// Config related
func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")

		flag.Parse()

		configPath = *flags


		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	_, err := os.Stat(configPath)
	if err != nil {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	var cfg Config

	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Cannot read config file: %s", err.Error())
	}

	return &cfg
}