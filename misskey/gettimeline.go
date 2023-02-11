package misskey

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/buger/jsonparser"
	"github.com/mattn/go-colorable"
)

var (
	Output = colorable.NewColorableStdout()
)

type noteData struct {
	offset    string
	timestamp string
	name      string
	username  string
	host      string
	text      string
	attach    string
	id        string
	isCat     bool
}

// タイムライン取得
func (c *Client) GetTimeline(limit int, mode string) error {
	body := struct {
		I     string `json:"i"`
		Limit int    `json:"limit"`
	}{
		I:     c.InstanceInfo.Token,
		Limit: limit,
	}

	var endpoint string
	if mode == "local" {
		endpoint = "notes/local-timeline"
	} else if mode == "global" {
		endpoint = "notes/global-timeline"
	} else if mode == "home" {
		endpoint = "notes/timeline"
	} else {
		return errors.New("Please select mode in local/home/global")
	}

	jsonByte, err := json.Marshal(body)
	if err != nil {
		return err
	}

	if err := c.apiPost(jsonByte, endpoint); err != nil {
		return err
	}

	fmt.Println("Timeline: " + mode + "  @" + c.InstanceInfo.UserName + " (" + c.InstanceInfo.Host + ")")
	printLine()

	jsonparser.ArrayEach(c.resBuf.Bytes(), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

		// とりあえずTextを持ってきてみる
		_, err = jsonparser.GetString(value, "renoteId")

		var note noteData

		if err != nil {
			note, err = pickNote(value)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			_, err = jsonparser.GetString(value, "replyId")
			if err == nil {
				replyParentValue, _, _, _ := jsonparser.Get(value, "reply")
				replyParent, err := pickNote(replyParentValue)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					return
				}
				repStr := fmt.Sprintf("%s \x1b[35m%s(@%s)\x1b[0m\t %s \x1b[32m%s\x1b[0m\x1b[34m(%s)\x1b[0m", replyParent.timestamp, replyParent.name, replyParent.username, replyParent.text, replyParent.attach, replyParent.id)
				fmt.Fprintln(Output, repStr)
				note.offset = "    "
			}

		} else { // renoteだったら

			renoteValue, _, _, _ := jsonparser.Get(value, "renote")

			note, err = pickNote(renoteValue)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			note.name = "[RN]" + note.name

		}

		str := fmt.Sprintf("%s%s \x1b[31m%s(@%s)\x1b[0m\t %s \x1b[32m%s\x1b[0m\x1b[34m(%s)\x1b[0m", note.offset, note.timestamp, note.name, note.username, note.text, note.attach, note.id)

		fmt.Fprintln(Output, str)
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	return nil
}

func pickNote(value []byte) (noteData, error) {
	var note noteData
	var err error
	// 投稿者
	note.name, _ = jsonparser.GetString(value, "user", "name")

	//投稿者ID
	note.username, err = jsonparser.GetString(value, "user", "username")
	if err != nil {
		return note, err
	}
	//ホスト名
	note.host, err = jsonparser.GetString(value, "user", "host")
	if err == nil {
		note.username = note.username + "@" + note.host
	}

	// 投稿時刻
	note.timestamp, err = jsonparser.GetString(value, "createdAt")
	if err != nil {
		return note, err
	}
	t, _ := time.ParseInLocation("2006-01-02T15:04:05Z", note.timestamp, time.UTC)
	note.timestamp = t.In(time.Local).Format("2006/01/02 15:04:05")

	// 本文
	note.text, _ = jsonparser.GetString(value, "text")

	//投稿ID(元投稿)
	note.id, err = jsonparser.GetString(value, "id")
	if err != nil {
		return note, err
	}

	//ねこかどうか
	isCat, err := jsonparser.GetBoolean(value, "user", "isCat")
	if isCat {
		note.name = note.name + "(Cat)"
	}

	// ファイルが有れば
	filesId, _, _, _ := jsonparser.Get(value, "files")
	if len(filesId) != 2 {
		note.attach = "   (添付有り)"
	}

	return note, nil
}
