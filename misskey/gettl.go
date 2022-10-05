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

func GetTl(configs config.Config, instanceName string, limit int) {
	var err error
	instanceInfo, err = getInstanceInfo(configs, instanceName)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := GetTlBody{
		I:     instanceInfo.Token,
		Limit: limit,
	}

	jsonByte, err := json.Marshal(body)

	resJsonByte, err := apiPost(jsonByte, "notes")
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonparser.ArrayEach(resJsonByte, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		name, _, _, _ := jsonparser.Get(value, "user", "name")
		nameStr := string(name)

		username, _, _, _ := jsonparser.Get(value, "user", "username")
		usernameStr := string(username)

		text, _, _, _ := jsonparser.Get(value, "text")
		textStr := string(text)

		id, _, _, _ := jsonparser.Get(value, "id")
		idStr := string(id)

		isCatByte, _, _, _ := jsonparser.Get(value, "user", "isCat")
		isCat := string(isCatByte)

		if isCat == "true" {
			nameStr = nameStr + "(Cat)"
		}

		str := fmt.Sprintf("\x1b[31m%s (@%s)\x1b[0m\t %s \t\x1b[34m(%s)\x1b[0m", nameStr, usernameStr, textStr, idStr)

		fmt.Println(str)
	})

}
