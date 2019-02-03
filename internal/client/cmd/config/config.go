package config

import (
	"github.com/BurntSushi/toml"
)

// Config for the commands
type Config struct {
	Port       string
	ConfigPath string
	Server     serverConfig
}

// Load the join config
func (c *Config) Load() error {
	if c.ConfigPath == "" {
		return c.setDefault()
	}

	var conf Config
	_, err := toml.DecodeFile(c.ConfigPath, &conf)
	if err != nil {
		return err
	}

	if c.Port == "" {
		c.Port = conf.Port
	}

	if err := c.Server.load(&conf.Server); err != nil {
		return err
	}

	return c.setDefault()
}

var (
	defaultPort = "9090"
)

func (c *Config) setDefault() error {
	if c.Port == "" {
		c.Port = defaultPort
	}
	return nil
}

type serverConfig struct {
	LibraryPath string
}

func (c *serverConfig) load(conf *serverConfig) error {
	if c.LibraryPath == "" {
		c.LibraryPath = conf.LibraryPath
	}

	return nil
}
