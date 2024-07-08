package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func Init() {
	CFG = MustLoad("configs/main.yml")
}

var CFG Config

type Config struct {
	HTTPServer   `yaml:"http_server"`
	Database     `yaml:"database"`
	JwtSecretKey string `yaml:"jwt_secret_key"`
}

type Database struct {
	URI string `yaml:"URI"`
}

type HTTPServer struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

func MustLoad(path string) Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("config files does not exist: %s", path)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("can't read config: %s", err)
	}

	return cfg
}
