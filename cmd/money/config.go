package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RunAddress  string `yaml:"RUN_ADDRESS"`
	WorkplaceDB string `yaml:"WORKPLACE_DB"`
	AuthDB      string `yaml:"AUTH_DB"`
}

func configuate() (*Config, error) {
	/*
		Приоритет настроек:
			1. Флаги;
			2. Конфигурационный файл
	*/

	flag.Parse()
	config, err := configFromFile()
	if err != nil {
		return nil, err
	}
	if *r != "" {
		config.RunAddress = *r
	}

	if *w != "" {
		config.WorkplaceDB = *w
	}

	if *a != "" {
		config.AuthDB = *a
	}

	return config, nil
}

func configFromFile() (*Config, error) {
	file, _ := os.Open("conf.yaml")
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Println("configFromFile:", err)
		return nil, err
	}

	fmt.Println(configuration)
	return &configuration, nil
}
