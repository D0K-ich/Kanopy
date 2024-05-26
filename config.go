package Kanopy

import (
	"errors"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/D0K-ich/Kanopy/logs"
)

type Config struct {
	Logger		*logs.Config		`yaml:"logger"`
}

func NewConfig(path *string) (config *Config, err error) {
	if *path = strings.TrimSpace(*path); *path == "" {err = errors.New("empty path for create main conf");return}

	var bytes []byte
	if bytes, err = os.ReadFile(*path); err != nil {return}

	if err = yaml.Unmarshal(bytes, &config); err != nil {return}
	return
}