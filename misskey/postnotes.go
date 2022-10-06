package misskey

import (
	"encoding/json"
	"fmt"

	"github.com/buger/jsonparser"
	"github.com/mikuta0407/misskey-cli/config"
)

// POST用
type CreateNoteBody struct {
	I    string `json:"i"`
	Text string `json:"text"`
}
type ReplyNoteBody struct {
	I       string `json:"i"`
	Text    string `json:"text"`
	ReplyId string `json:"replyId"`
}

type RenoteNoteBody struct {
	I        string `json:"i"`
	RenoteId string `json:"renoteId"`
}

type DeleteNoteBody struct {
	I      string `json:"i"`
	NoteId string `json:"noteId"`
}

func CreateNote(configs config.Config, instanceName string, text string) {
	var err error
	instanceInfo, err = getInstanceInfo(configs, instanceName)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := CreateNoteBody{
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

func ReplyNote(configs config.Config, instanceName string, replyId string, text string) {
	var err error
	instanceInfo, err = getInstanceInfo(configs, instanceName)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := ReplyNoteBody{
		I:       instanceInfo.Token,
		Text:    text,
		ReplyId: replyId,
	}

	jsonByte, err := json.Marshal(body)

	resJsonByte, err := apiPost(jsonByte, "notes/create")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(resJsonByte))
	// id, _ := jsonparser.GetString(resJsonByte, "createdNote", "id")
	// text, _ = jsonparser.GetString(resJsonByte, "createdNote", "text")

	// str := fmt.Sprintf("Note Success! id : %s\n\"%s\"", string(id), string(text))

	// fmt.Println(str)

}

func RenoteNote(configs config.Config, instanceName string, renoteId string) {
	var err error
	instanceInfo, err = getInstanceInfo(configs, instanceName)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := RenoteNoteBody{
		I:        instanceInfo.Token,
		RenoteId: renoteId,
	}

	jsonByte, err := json.Marshal(body)

	resJsonByte, err := apiPost(jsonByte, "notes/create")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(resJsonByte))

	// id, _ := jsonparser.GetString(resJsonByte, "createdNote", "id")
	// text, _ = jsonparser.GetString(resJsonByte, "createdNote", "text")

	// str := fmt.Sprintf("Note Success! id : %s\n\"%s\"", string(id), string(text))

	// fmt.Println(str)

}

func DeleteNote(configs config.Config, instanceName string, noteId string) {
	var err error
	instanceInfo, err = getInstanceInfo(configs, instanceName)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := DeleteNoteBody{
		I:      instanceInfo.Token,
		NoteId: noteId,
	}

	jsonByte, err := json.Marshal(body)

	resJsonByte, err := apiPost(jsonByte, "notes/delete")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Deleted!")

	fmt.Println(string(resJsonByte))

	// id, _ := jsonparser.GetString(resJsonByte, "createdNote", "id")
	// text, _ = jsonparser.GetString(resJsonByte, "createdNote", "text")

	// str := fmt.Sprintf("Note Success! id : %s\n\"%s\"", string(id), string(text))

	// fmt.Println(str)

}
