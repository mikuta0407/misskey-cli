package misskey

import (
	"encoding/json"
	"fmt"

	"github.com/buger/jsonparser"
)

// POST用
type CreateNoteBody struct {
	I    string `json:"i"`
	Text string `json:"text"`
}

type NoteBody struct {
	I      string `json:"i"`
	NoteId string `json:"noteId"`
}

func (c *Client) CreateNote(text string) error {
	body := struct {
		I    string `json:"i"`
		Text string `json:"text"`
	}{
		I:    c.InstanceInfo.Token,
		Text: text,
	}

	jsonByte, err := json.Marshal(body)
	if err != nil {
		return err
	}

	if err := c.apiPost(jsonByte, "notes/create"); err != nil {
		return err
	}

	id, _ := jsonparser.GetString(c.resBuf.Bytes(), "createdNote", "id")
	text, _ = jsonparser.GetString(c.resBuf.Bytes(), "createdNote", "text")

	str := fmt.Sprintf("Note Success! id : %s\n\"%s\"", string(id), string(text))

	fmt.Println(str)

	return nil

}

func (c *Client) ReplyNote(replyId string, text string) error {
	body := struct {
		I       string `json:"i"`
		ReplyId string `json:"replyId"`
		Text    string `json:"text"`
	}{
		I:       c.InstanceInfo.Token,
		ReplyId: replyId,
		Text:    text,
	}

	jsonByte, err := json.Marshal(body)
	if err != nil {
		return err
	}

	if err := c.apiPost(jsonByte, "notes/create"); err != nil {
		return err
	}

	fmt.Println(c.resBuf)

	return nil
	// id, _ := jsonparser.GetString(resJsonByte, "createdNote", "id")
	// text, _ = jsonparser.GetString(resJsonByte, "createdNote", "text")

	// str := fmt.Sprintf("Note Success! id : %s\n\"%s\"", string(id), string(text))

	// fmt.Println(str)

}

func (c *Client) RenoteNote(renoteId string) error {
	body := struct {
		I        string `json:"i"`
		RenoteId string `json:"replyId"`
	}{
		I:        c.InstanceInfo.Token,
		RenoteId: renoteId,
	}

	jsonByte, err := json.Marshal(body)
	if err != nil {
		return err
	}

	if err := c.apiPost(jsonByte, "notes/create"); err != nil {
		return err
	}

	fmt.Println(c.resBuf)
	return nil

	// id, _ := jsonparser.GetString(resJsonByte, "createdNote", "id")
	// text, _ = jsonparser.GetString(resJsonByte, "createdNote", "text")

	// str := fmt.Sprintf("Note Success! id : %s\n\"%s\"", string(id), string(text))

	// fmt.Println(str)

}

func (c *Client) DeleteNote(noteId string) error {

	body := struct {
		I      string `json:"i"`
		NoteId string `json:"noteId"`
	}{
		I:      c.InstanceInfo.Token,
		NoteId: noteId,
	}

	jsonByte, err := json.Marshal(body)
	if err != nil {
		return err
	}

	if err := c.apiPost(jsonByte, "notes/delete"); err != nil {
		return err
	}

	fmt.Println("Deleted!")

	return nil

	// id, _ := jsonparser.GetString(resJsonByte, "createdNote", "id")
	// text, _ = jsonparser.GetString(resJsonByte, "createdNote", "text")

	// str := fmt.Sprintf("Note Success! id : %s\n\"%s\"", string(id), string(text))

	// fmt.Println(str)

}
