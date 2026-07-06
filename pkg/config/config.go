package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	Data struct {
		LoadToRamOnInit		bool `toml:"load_to_ram_on_init"`
		Path	string `toml:"path"`
	} `toml:"data"`
}

func GetConfigDirectory() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, OROFACIAL_ATLAS_DIRNAME, "config.toml"), nil
}

func DoesConfigExist() bool {
	configDir, err := GetConfigDirectory()
	if err != nil {
		return false
	}

	_, err = os.Stat(configDir)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func GetConfigs() (*Config, error) {
	return nil, nil
}
