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

		var (
			nameStr     string
			usernameStr string
			textStr     string
			attach      string
			idStr       string
		)

		text, _, _, _ := jsonparser.Get(value, "text")

		if string(text) != "null" {
			// renoteじゃなかったら
			// 投稿者
			name, _, _, _ := jsonparser.Get(value, "user", "name")
			nameStr = string(name)

			//投稿者ID
			username, _, _, _ := jsonparser.Get(value, "user", "username")
			usernameStr = string(username)

			//ホスト名
			host, _, _, _ := jsonparser.Get(value, "user", "host")
			if string(host) != "null" {
				usernameStr = usernameStr + "@" + string(host)
			}

			// 本文
			text, _, _, _ := jsonparser.Get(value, "text")
			textStr = string(text)

			//投稿ID(元投稿)
			id, _, _, _ := jsonparser.Get(value, "id")
			idStr = string(id)

			//ねこかどうか
			isCatByte, _, _, _ := jsonparser.Get(value, "user", "isCat")
			isCat := string(isCatByte)

			if isCat == "true" {
				nameStr = nameStr + "(Cat)"
			}

			// ファイルが有れば
			filesId, _, _, _ := jsonparser.Get(value, "files")
			if len(filesId) != 2 {
				attach = "   (添付有り)"
			}
		} else {
			// renoteだったら
			// 投稿者
			name, _, _, _ := jsonparser.Get(value, "renote", "user", "name")
			nameStr = "[RN]" + string(name)

			//投稿者ID
			username, _, _, _ := jsonparser.Get(value, "renote", "user", "username")
			usernameStr = string(username)

			//ホスト名
			host, _, _, _ := jsonparser.Get(value, "renote", "user", "host")
			if string(host) != "null" {
				usernameStr = usernameStr + "@" + string(host)
			}

			// 本文
			text, _, _, _ := jsonparser.Get(value, "renote", "text")
			textStr = string(text)

			//投稿ID(元投稿)
			id, _, _, _ := jsonparser.Get(value, "renote", "id")
			idStr = string(id)

			//ねこかどうか
			isCatByte, _, _, _ := jsonparser.Get(value, "renote", "user", "isCat")
			isCat := string(isCatByte)

			if isCat == "true" {
				nameStr = nameStr + "(Cat)"
			}

			// ファイルが有れば
			filesId, _, _, _ := jsonparser.Get(value, "renote", "files")
			if len(filesId) != 2 {
				attach = "   (添付有り)"
			}
		}

		str := fmt.Sprintf("\x1b[31m%s (@%s)\x1b[0m\t %s \t\x1b[32m%s\x1b[0m\x1b[34m(%s)\x1b[0m", nameStr, usernameStr, textStr, attach, idStr)

		fmt.Println(str)
	})

}
