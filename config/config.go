package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// New creates a Config from a JSON file.
func New(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("config: opening client config file: %s", err)
	}

	var conf Config
	if err = json.Unmarshal(file, &conf); err != nil {
		return nil, fmt.Errorf("config: unmarshal config json: %s", err)
	}
	conf.loaded = true
	return &conf, nil
}

// Config is Crypt client config.
type Config struct {
	Key string
	URL string

	loaded bool
}

// Loaded verifies if the struct was created by the LoadConfig struct.
func (c *Config) Loaded() bool {
	return c.loaded
}
