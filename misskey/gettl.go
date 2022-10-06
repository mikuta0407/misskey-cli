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
	Local bool   `json:"local"`
}

// 以下受け取り用

func GetTl(configs config.Config, instanceName string, limit int, mode string) {
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

	var endpoint string
	if mode == "local" {
		endpoint = "notes/local-timeline"
	} else if mode == "global" {
		endpoint = "notes/global-timeline"
	} else if mode == "home" {
		endpoint = "notes/timeline"
	} else {
		fmt.Println("Please select mode in local/home/global")
		return
	}
	resJsonByte, err := apiPost(jsonByte, endpoint)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonparser.ArrayEach(resJsonByte, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		name, _, _, _ := jsonparser.Get(value, "user", "name")
		nameStr := string(name)

		username, _, _, _ := jsonparser.Get(value, "user", "username")
		usernameStr := string(username)

		host, _, _, _ := jsonparser.Get(value, "user", "host")
		if string(host) != "null" {
			usernameStr = usernameStr + "@" + string(host)
		}

		text, _, _, _ := jsonparser.Get(value, "text")
		textStr := string(text)

		id, _, _, _ := jsonparser.Get(value, "id")
		idStr := string(id)

		isCatByte, _, _, _ := jsonparser.Get(value, "user", "isCat")
		isCat := string(isCatByte)

		if isCat == "true" {
			nameStr = nameStr + "(Cat)"
		}

		filesId, _, _, _ := jsonparser.Get(value, "files")
		var attach string
		if len(filesId) != 2 {
			attach = "   (添付有り)"
		}

		str := fmt.Sprintf("\x1b[31m%s (@%s)\x1b[0m\t %s \t\x1b[32m%s\x1b[0m\x1b[34m(%s)\x1b[0m", nameStr, usernameStr, textStr, attach, idStr)

		fmt.Println(str)
	})

}
