package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Settings struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func GetConfig(path string) *Settings {

	var config Settings

	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := yaml.NewDecoder(f).Decode(&config); err != nil {
		log.Fatalln(err.Error())
	}

	return &config
}
