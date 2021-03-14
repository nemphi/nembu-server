package config

import (
	"encoding/json"
)

// Config for this instance of nembu-server
type Config struct {
	DB DBConfig `json:"db,omitempty"`
}

// DBConfig to use for this instance of nembu-server
type DBConfig struct {
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
}

func LoadFromJSON(bytes []byte) (*Config, error) {
	cfg := &Config{}
	err := json.Unmarshal(bytes, cfg)
	return cfg, err
}
