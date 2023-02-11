package misskey

import (
	"bytes"
	"fmt"
	"os"
	"syscall"

	"github.com/mikuta0407/misskey-cli/config"
	"golang.org/x/crypto/ssh/terminal"
)

type Client struct {
	InstanceInfo config.InstanceInfo
	resBuf       *bytes.Buffer
}

func NewClient(instanceName string, cfgFile string) *Client {
	configs, err := config.ParseToml(cfgFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var instanceInfo config.InstanceInfo

	if instanceName != "" {
		index, isExist := include(configs.Instance, instanceName)

		if isExist {
			instanceInfo = configs.Instance[index]
		} else {
			fmt.Fprintln(os.Stderr, "No instance name in config")
			os.Exit(1)
		}
	} else {
		instanceInfo = configs.Instance[0]
	}

	if instanceInfo.Host == "" || instanceInfo.Token == "" {
		fmt.Fprintln(os.Stderr, "No instance specification")
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

func printLine() {
	width, _, err := terminal.GetSize(syscall.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %+v", err)
		os.Exit(1)
	}

	for i := 1; i <= width; i++ {
		fmt.Printf("=")
	}
	fmt.Printf("\n")
}
