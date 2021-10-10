package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

var AppConfig config

type config struct {
	MovieService movieService `toml:"movie_service"`
	GRPC         grpc         `toml:"grpc"`
}

type movieService struct {
	BaseUrl      string   `toml:"base_url"`
	SecretKey    string   `toml:"secret_key"`
	TimeoutInSec duration `toml:"timeout_in_sec"`
}

type grpc struct {
	Port string `toml:"port"`
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func InitializeAppConfig() {
	if _, err := toml.DecodeFile(configPath(), &AppConfig); err != nil {
		panic(err)
	}
}

func configPath() string {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(workPath, "../config", "config.toml")
}
