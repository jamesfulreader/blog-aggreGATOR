package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = "gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name,omitempty"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}

func Read() (Config, error) {
	config, err := os.ReadFile(getConfigFilePath())
	if err != nil {
		return Config{}, err
	}

	var data Config

	if err := json.Unmarshal(config, &data); err != nil {
		return Config{}, fmt.Errorf("error unmarshaling due to %s", err)
	}

	return data, nil
}

func getConfigFilePath() string {
	projectRoot := filepath.Join("cd ./../../gator", configFileName)
	return projectRoot
}

func write(c Config) error {
	file, err := os.Create(getConfigFilePath())
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		return err
	}

	return nil
}
