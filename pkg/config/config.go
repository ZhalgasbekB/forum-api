package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port    string `json:"port"`
	Db      struct {
		Driver string `json:"driver"`
		Dsn    string `json:"dsn"`
	} `json:"db"`
}

func Load() (Config, error) {
	var cfg Config
	file, err := os.Open("./config/config.json")
	if err != nil {
		return Config{}, fmt.Errorf("Конфигурация не загружена: %w", err)
	}
	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("Конфигурация не конвертирована: %w", err)
	}
	return cfg, nil
}
