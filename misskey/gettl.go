package misskey

import (
	"encoding/json"
	"fmt"

	"github.com/mikuta0407/misskey-cli/config"
)

func GetTl(configs config.Config, instanceName string) {
	var err error
	instanceInfo, err = getInstanceInfo(configs, instanceName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(instanceInfo)

	var body struct {
		i     string `json:"i"`
		limit int    `json:"limit"`
	}
	body.i = instanceInfo.Token
	body.limit = 10
	jsonByte, err := json.Marshal(body)
	fmt.Println(body)
	fmt.Println(json.Marshal(body))
	apiPost(jsonByte, "notes")

}
