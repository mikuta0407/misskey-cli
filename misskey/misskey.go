package misskey

import (
	"bytes"
	"fmt"
	"os"

	"github.com/mikuta0407/misskey-cli/config"
)

type Client struct {
	InstanceInfo config.InstanceInfo
	resBuf       *bytes.Buffer
}

func NewClient(instanceName string, cfgFile string) *Client {
	configs, err := config.ParseToml(cfgFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var instanceInfo config.InstanceInfo
	index, isExist := include(configs.Instance, instanceName)
	if isExist {
		instanceInfo = configs.Instance[index]
	} else {
		fmt.Println("No instance name in config")
		return nil
	}

	return &Client{
		InstanceInfo: instanceInfo,
	}
}

func include(slice []config.InstanceInfo, target string) (int, bool) {
	for i, info := range slice {
		if info.Name == target {
			return i, true
		}
	}
	return -1, false
}
