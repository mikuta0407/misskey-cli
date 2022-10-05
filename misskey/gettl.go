package misskey

import (
	"encoding/json"
	"fmt"

	"github.com/buger/jsonparser"
	"github.com/mikuta0407/misskey-cli/config"
)

// POST用
type GetTlBody struct {
	I     string `json:"i"`
	Limit int    `json:"limit"`
}

// 以下受け取り用

func GetTl(configs config.Config, instanceName string) {
	var err error
	instanceInfo, err = getInstanceInfo(configs, instanceName)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := GetTlBody{
		I:     instanceInfo.Token,
		Limit: 10,
	}

	jsonByte, err := json.Marshal(body)

	resJsonByte, err := apiPost(jsonByte, "notes")
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonparser.ArrayEach(resJsonByte, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		name, _, _, _ := jsonparser.Get(value, "user", "name")

		username, _, _, _ := jsonparser.Get(value, "user", "username")

		text, _, _, _ := jsonparser.Get(value, "text")

		id, _, _, _ := jsonparser.Get(value, "id")

		str := fmt.Sprintf("\x1b[31m%s (@%s)\x1b[0m\t %s \t\x1b[34m(%s)\x1b[0m", string(name), string(username), string(text), string(id))

		fmt.Println(str)
	})

}
