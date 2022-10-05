package misskey

import (
	"encoding/json"
	"fmt"

	"github.com/mikuta0407/misskey-cli/config"
)

// POST用
type PostNoteBody struct {
	I    string `json:"i"`
	Text string `json:"text"`
}

func PostNote(configs config.Config, instanceName string, text string) {
	var err error
	instanceInfo, err = getInstanceInfo(configs, instanceName)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := PostNoteBody{
		I:    instanceInfo.Token,
		Text: text,
	}

	jsonByte, err := json.Marshal(body)

	resJsonByte, err := apiPost(jsonByte, "notes/create")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(resJsonByte))

	// jsonparser.ArrayEach(resJsonByte, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
	// 	name, _, _, _ := jsonparser.Get(value, "user", "name")

	// 	username, _, _, _ := jsonparser.Get(value, "user", "username")

	// 	text, _, _, _ := jsonparser.Get(value, "text")

	// 	id, _, _, _ := jsonparser.Get(value, "id")

	// 	str := fmt.Sprintf("\x1b[31m%s (@%s)\x1b[0m\t %s \t\x1b[34m(%s)\x1b[0m", string(name), string(username), string(text), string(id))

	// 	fmt.Println(str)
	// })

}
