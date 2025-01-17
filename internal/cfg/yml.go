package cfg

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
