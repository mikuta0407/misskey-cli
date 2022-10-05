package config


import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Instance []InstanceInfo `toml:"Instance"`
}

type InstanceInfo struct {
	Host     string `toml:"host" validate:"required"`
	Name     string `toml:"name" validate:"required"`
	Token    string `toml:"token" validate:"required"`
	UserName string `toml:"username" validate:"required"`
}

func ParseToml(fileName string) (Config, error) {
	fmt.Println(fileName)
	var configs Config
	_, err := toml.DecodeFile(fileName, &configs)

	return configs, err
}
