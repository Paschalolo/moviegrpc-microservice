package main

type serviceConfig struct {
	APIConfig apiConfig    `yaml:"api"`
	Jaeger    jaegerConfig `yaml:"jaeger"`
}
type apiConfig struct {
	Port int `yaml:"port"`
}

type jaegerConfig struct {
	URL string `yaml:"url"`
}
