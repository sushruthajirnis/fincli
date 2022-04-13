package config

import (
	"errors"
	"github.com/spf13/viper"
	"os"
	"os/user"
	"path/filepath"
)

//https://github.com/datreeio/datree
// heavily influenced by ppl at datree

const (
	tokenKey = "token"
)

type LocalConfig struct {
	Token string
}

type LocalConfigClient struct {
}

func NewLocalConfigClient() *LocalConfigClient {
	return &LocalConfigClient{}
}

func (lc *LocalConfigClient) GetLocalConfiguration() (*LocalConfig, error) {
	viper.SetEnvPrefix("fincli")
	viper.AutomaticEnv()
	initConfigFileErr := InitLocalConfigFile()
	if initConfigFileErr != nil {
		return nil, initConfigFileErr
	}
	token := viper.GetString(tokenKey)

	if token == "" {
		return nil, errors.New("unable to find api token")

	}
	return &LocalConfig{
		Token: token,
	}, nil
}

func InitLocalConfigFile() error {
	configHome, configName, configType, err := setViperConfig()
	if err != nil {
		return err
	}
	// workaround for creating config file when not exist
	// open issue in viper: https://github.com/spf13/viper/issues/430
	// should be fixed in pr https://github.com/spf13/viper/pull/936
	configPath := filepath.Join(configHome, configName+"."+configType)

	// workaround for catching error if not enough permissions
	// resolves issues in https://github.com/Homebrew/homebrew-core/pull/97061
	isConfigExists, _ := exists(configPath)
	if !isConfigExists {
		_ = os.Mkdir(configHome, os.ModePerm)
		_, _ = os.Create(configPath)
	}

	_ = viper.ReadInConfig()
	return nil
}
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// config example from datree github.com/datreeio/datree
func getConfigHome() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	homedir := usr.HomeDir
	configHome := filepath.Join(homedir, ".fincli")

	return configHome, nil
}

func getConfigName() string {
	return "config"
}
func getConfigType() string {
	return "yaml"
}
func setViperConfig() (string, string, string, error) {
	configHome, err := getConfigHome()
	if err != nil {
		return "", "", "", nil
	}

	configName := getConfigName()
	configType := getConfigType()

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configHome)

	return configHome, configName, configType, nil
}

func (lc *LocalConfigClient) Set(key string, value string) error {
	initConfigFileErr := InitLocalConfigFile()
	if initConfigFileErr != nil {
		return initConfigFileErr
	}

	viper.Set(key, value)
	writeClientIdErr := viper.WriteConfig()
	if writeClientIdErr != nil {
		return writeClientIdErr
	}
	return nil
}
