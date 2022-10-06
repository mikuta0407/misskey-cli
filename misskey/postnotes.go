package misskey

import (
	"encoding/json"
	"fmt"

	"github.com/buger/jsonparser"
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

	id, _ := jsonparser.GetString(resJsonByte, "createdNote", "id")
	text, _ = jsonparser.GetString(resJsonByte, "createdNote", "text")

	str := fmt.Sprintf("Note Success! id : %s\n\"%s\"", string(id), string(text))

	fmt.Println(str)

}
