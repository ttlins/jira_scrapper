package config

import (
	"log"
	"path"

	"github.com/go-playground/validator"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const appName = "jira_scrapper"

var (
	home = homePath()
	cfg  Configuration
)

type Project struct {
	Key     string `yaml:"key" validate:"required"`
	BoardID int    `yaml:"boardId"`
}

type Configuration struct {
	Jira struct {
		Token    string    `yaml:"token" validate:"required"`
		Host     string    `yaml:"host" validate:"required"`
		Projects []Project `yaml:"projects"`
	} `yaml:"jira"`
}

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(Path())
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}
	validateConfig()
}

func Config() Configuration {
	return cfg
}

func Path() string {
	return path.Join(home, "."+appName)
}

func validateConfig() {
	validate := validator.New()
	if err := validate.Struct(&cfg); err != nil {
		log.Fatalf("failed to validate config: %v", err)
	}
	for _, p := range cfg.Jira.Projects {
		if err := validate.Struct(&p); err != nil {
			log.Fatalf("failed to validate config: %v", err)
		}
	}
}

func homePath() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("failed to get home dir: %v", err)
	}
	return home
}
