package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	route = "config/config.yaml"
)

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Storage struct {
	APIURLBase   string `yaml:"api-url-base"`
	ProjectName  string `yaml:"project-name"`
	DatabaseName string `yaml:"database-name"`
	DatabaseKey  string `yaml:"database-key"`
}

type Config struct {
	Service Service `yaml:"data-service-config"`
	Storage Storage `yaml:"cloud-storage-config"`
}

func Initialise() (*Config, error) {

	conf := Config{}

	yamlFile, err := ioutil.ReadFile(route)
	if err != nil {
		return &Config{}, fmt.Errorf("issue finding config yaml, err %v", err)
	}
	yamlFile = []byte(os.ExpandEnv(string(yamlFile)))

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return &Config{}, fmt.Errorf("issue unmarshalling config yaml, err %v", err)
	}

	return &conf, nil
}
