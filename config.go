package main

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/Masterminds/semver/v3"
	"gopkg.in/yaml.v3"
)

const (
	IncrementMajor = "major"
	IncrementMinor = "minor"
	IncrementPatch = "patch"
)

type Config struct {
	Increment               string          `yaml:"increment"`
	Verify                  string          `yaml:"verify"`
	Perform                 string          `yaml:"perform"`
	Version                 *semver.Version `yaml:"-"`
	SkipPush                bool            `yaml:"skip-push"`
	SkipCleanWorkspaceCheck bool            `yaml:"skip-clean-workspace-check"`
}

func LoadConfig() (*Config, error) {

	config := &Config{
		Increment: IncrementMinor,
	}

	y, err := os.ReadFile(path.Join(".", ".go-release.yaml"))

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, errors.New("config file .go-release.yaml does not exist")
		}
		return nil, err
	}

	err = yaml.Unmarshal(y, config)

	if err != nil {
		return nil, err
	}

	v, err := os.ReadFile(path.Join(".", ".version"))

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, errors.New("version file .version does not exist")
		}
		return nil, err
	}

	versionString := strings.TrimSpace(string(v))

	if strings.Contains(versionString, "\n") {
		return nil, errors.New(".version file must be on one line")
	}

	version, err := semver.NewVersion(versionString)

	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(version.Original(), "v") {
		return nil, errors.New("version must start with 'v'")
	}

	config.Version = version

	return config, nil
}
