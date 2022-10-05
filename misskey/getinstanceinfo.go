package misskey

import (
	"errors"

	"github.com/mikuta0407/misskey-cli/config"
)

func include(slice []config.InstanceInfo, target string) (int, bool) {
	for i, info := range slice {
		if info.Name == target {
			return i, true
		}
	}
	return -1, false
}

func getInstanceInfo(configs config.Config, instanceName string) (config.InstanceInfo, error) {
	index, isExist := include(configs.Instance, instanceName)
	if isExist {
		return configs.Instance[index], nil
	} else {
		return configs.Instance[0], errors.New("No instance name in config")
	}
}

var instanceInfo config.InstanceInfo
