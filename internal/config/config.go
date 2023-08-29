package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	dbPassEscSeq = "{password}"
	password     = "avito_pass"
)

type Grpc struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	Host                 string `yaml:"host"`
	Port                 string `yaml:"port"`
	User                 string `yaml:"user"`
	Name                 string `yaml:"database"`
	Ssl                  string `yaml:"ssl"`
	MaxOpenedConnections int32  `yaml:"max_opened_connections"`
}

type Http struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Config struct {
	Grpc     Grpc     `yaml:"grpc"`
	Database Database `yaml:"database"`
	Http     Http     `yaml:"http"`
}

func Read(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) GetDBConfig() (string, error) {
	DbDsn := fmt.Sprintf("user=%s dbname=%s password={password} host=%s port=%s sslmode=%s", c.Database.User, c.Database.Name, c.Database.Host, c.Database.Port, c.Database.Ssl)
	DbDsn = strings.ReplaceAll(DbDsn, dbPassEscSeq, password)

	return DbDsn, nil
}
