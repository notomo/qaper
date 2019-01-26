package config

import "github.com/BurntSushi/toml"

// Config for the commands
type Config struct {
	Server ServerConfig
}

// ServerConfig configs `server` command
type ServerConfig struct {
	LibraryPath string
	Port        string
}

// Load the server config
func (c *ServerConfig) Load(configPath string) (*ServerConfig, error) {
	if configPath == "" {
		return c, nil
	}

	var conf Config
	_, err := toml.DecodeFile(configPath, &conf)
	if err != nil {
		return nil, err
	}

	serverConfig := conf.Server
	if c.Port != "" {
		serverConfig.Port = c.Port
	}
	if c.LibraryPath != "" {
		serverConfig.LibraryPath = c.LibraryPath
	}

	return &serverConfig, nil
}
