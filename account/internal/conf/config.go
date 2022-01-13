package conf

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func InitConfig(configFilePath string) (Configuration, error) {
	f, err := os.Open(configFilePath)
	if err != nil {
		return Configuration{}, errors.Wrap(err, "unable to open configuration file")
	}
	defer func() { _ = f.Close() }()

	var config = &Configuration{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)

	if err != nil {
		return Configuration{}, errors.Wrap(err, "unable to load Configuration")
	}
	return *config, nil
}
