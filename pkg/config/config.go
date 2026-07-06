package config

import (
	"os"
	"path/filepath"

	"github.com/brandondkong/orofacial-atlas/pkg/input"
	"github.com/pelletier/go-toml/v2"
)

var cfg *Config = nil

type DataConfig struct {
	LoadToRamOnInit		bool `toml:"load_to_ram_on_init"`
	Path	string `toml:"path"`
}

type Config struct {
	EnableAiFeatures	bool `toml:"enable_ai_features"`
	Data DataConfig `toml:"data"`
}

func getConfigDirectory() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, OROFACIAL_ATLAS_DIRNAME), nil
}

func createConfigDirectory() error {
	cfgDir, err := getConfigDirectory()
	if err != nil {
		return err
	}

	return os.MkdirAll(cfgDir, 0777)
}

func getConfigFilePath() (string, error) {
	cfgDir, err := getConfigDirectory()
	if err != nil {
		return "", err
	}

	return filepath.Join(cfgDir, "config.toml"), nil
}


func getDefaultConfigDataDir() (string, error) {
	cfgDir, err := getConfigDirectory()
	if err != nil {
		return "", err
	}

	return filepath.Join(cfgDir, DEFAULT_DATA_DIRECTORY), nil
}

func DoesConfigExist() bool {
	configDir, err := getConfigFilePath()
	if err != nil {
		return false
	}

	_, err = os.Stat(configDir)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func GetConfigs() *Config {
	return cfg
}

func PromptNewConfigs() error {
	enableAi, err := input.PromptYesNoAnswer("Would you like to enable AI features? (y/n): ")
	if err != nil {
		return err
	}

	dataPath, err := input.PromptString("Which path would you like data to be stored? (default: ~/.orofacial-atlas/data): ")
	if err != nil {
		return err
	}

	if dataPath == "" {
		dataPath, err = getDefaultConfigDataDir()
		if err != nil {
			return err
		}
	}

	cfg = &Config{
		EnableAiFeatures: enableAi,
		Data: DataConfig{
			Path: dataPath,
		},
	}

	fd, err := toml.Marshal(cfg)
	if err != nil {
		return err
	}

	cfgPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// create the directory if it doesn't exist
	err = createConfigDirectory()
	if err != nil {
		return err
	}

	err = os.WriteFile(cfgPath, fd, 0644)
	return err
}
