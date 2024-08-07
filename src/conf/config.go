package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	SQL    SQLConfig    `yaml:"sql"`
	Redis  RedisConfig  `yaml:"redis"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Gin  struct {
		Mode string `yaml:"mode"`
		CORS struct {
			Enabled          bool     `yaml:"enabled"`
			AllowedOrigins   []string `yaml:"allowed_origins"`
			AllowedMethods   []string `yaml:"allowed_methods"`
			AllowHeaders     []string `yaml:"allowed_headers"`
			ExposedHeaders   []string `yaml:"exposed_headers"`
			AllowCredentials bool     `yaml:"allow_credentials"`
		} `yaml:"cors"`
	} `yaml:"gin"`
	SessionCookie struct {
		Name     string `yaml:"name"`
		Domain   string `yaml:"domain"`
		MaxAge   int    `yaml:"max_age"`
		Secure   bool   `yaml:"secure"`
		HttpOnly bool   `yaml:"http_only"`
	} `yaml:"session_cookie"`
}

type SQLConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Driver   string `yaml:"driver"`
	Database string `yaml:"database"`
	Protocol string `yaml:"protocol"`
}

type RedisConfig struct {
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

func (co *Config) GetConfig() Config {
	config := Config{}

	// Open config file
	filename := flag.String("config", "./src/conf/conf.yaml", "Location of the config file.")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		panic(err)
	}

	return config
}
