package main

import (
	"flag"
	"os"
)

var (
	a, r, w *string
)

func init() {
	r = flag.String("r", "", "RUN_ADDRESS")
	w = flag.String("w", "", "WORKPLACE_DB")
	a = flag.String("a", "", "AUTH_DB")
}

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

	config := &Config{RunAddress: os.Getenv("RUN_ADDRESS"), WorkplaceDB: os.Getenv("DATABASE_URI"), AuthDB: os.Getenv("DATABASE_URI")}

	flag.Parse()
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

/* obsolete
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
*/
