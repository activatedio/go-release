package main

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Config struct {
	Verify  string
	Perform string
	Version *Version
}

func LoadConfig() (*Config, error) {

	config := &Config{}

	y, err := os.ReadFile("./go-release.yaml")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, errors.New("config file ./go-release.yaml does not exist")
		} else {
			return nil, err
		}
	}

	err = yaml.Unmarshal(y, config)

	if err != nil {
		return nil, err
	}

	v, err := os.ReadFile("./.version")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, errors.New("version file ./.version does not exist")
		} else {
			return nil, err
		}
	}

	versionString := strings.TrimSpace(string(v))

	if strings.Contains(versionString, "\n") {
		return nil, errors.New(".version file must be on one line")
	}

	version, err := ParseVersion(versionString)

	if err != nil {
		return nil, err
	}

	config.Version = version

	return config, nil
}
