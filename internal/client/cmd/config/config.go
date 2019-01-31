package config

import (
	"github.com/BurntSushi/toml"
)

// Config for the commands
type Config struct {
	Port     string
	Server   ServerConfig
	Join     JoinConfig
	Question QuestionConfig
	Answer   AnswerConfig
}

var defaultPort = "9090"

// ServerConfig configs `server` command
type ServerConfig struct {
	LibraryPath string
	Port        string
}

// Load the server config
func (c *ServerConfig) Load(configPath string) (*ServerConfig, error) {
	if configPath == "" {
		return c.setDefault(), nil
	}

	var conf Config
	_, err := toml.DecodeFile(configPath, &conf)
	if err != nil {
		return nil, err
	}

	serverConfig := conf.Server
	if c.Port == "" {
		c.Port = serverConfig.Port
	}
	if c.LibraryPath == "" {
		c.LibraryPath = serverConfig.LibraryPath
	}

	return c.setDefault(), nil
}

func (c *ServerConfig) setDefault() *ServerConfig {
	if c.Port == "" {
		c.Port = defaultPort
	}
	return c
}

// JoinConfig configs `join` command
type JoinConfig struct {
	Port string
}

// Load the join config
func (c *JoinConfig) Load(configPath string) (*JoinConfig, error) {
	if configPath == "" {
		return c.setDefault(), nil
	}

	var conf Config
	_, err := toml.DecodeFile(configPath, &conf)
	if err != nil {
		return nil, err
	}

	joinConfig := conf.Join
	if c.Port != "" {
		joinConfig.Port = c.Port
	}

	return c.setDefault(), nil
}

func (c *JoinConfig) setDefault() *JoinConfig {
	if c.Port == "" {
		c.Port = defaultPort
	}
	return c
}

// QuestionConfig configs `question` command
type QuestionConfig struct {
	Port string
}

// Load the question config
func (c *QuestionConfig) Load(configPath string) (*QuestionConfig, error) {
	if configPath == "" {
		return c.setDefault(), nil
	}

	var conf Config
	_, err := toml.DecodeFile(configPath, &conf)
	if err != nil {
		return nil, err
	}

	questionConfig := conf.Question
	if c.Port != "" {
		questionConfig.Port = c.Port
	}

	return c.setDefault(), nil
}

func (c *QuestionConfig) setDefault() *QuestionConfig {
	if c.Port == "" {
		c.Port = defaultPort
	}
	return c
}

// AnswerConfig configs `answer` command
type AnswerConfig struct {
	Port string
}

// Load the answer config
func (c *AnswerConfig) Load(configPath string) (*AnswerConfig, error) {
	if configPath == "" {
		return c.setDefault(), nil
	}

	var conf Config
	_, err := toml.DecodeFile(configPath, &conf)
	if err != nil {
		return nil, err
	}

	answerConfig := conf.Answer
	if c.Port != "" {
		answerConfig.Port = c.Port
	}

	return c.setDefault(), nil
}

func (c *AnswerConfig) setDefault() *AnswerConfig {
	if c.Port == "" {
		c.Port = defaultPort
	}
	return c
}
