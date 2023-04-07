package misskey

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/buger/jsonparser"
	"github.com/google/uuid"
	"github.com/sacOO7/gowebsocket"
)

// gettimeline.goに依存した実装

func (c *Client) GetStream(mode string) error {

	fmt.Println("Stream: " + mode + "  @" + c.InstanceInfo.UserName + " (" + c.InstanceInfo.Host + ")")
	printLine()

	parsedUrl, err := url.Parse(c.InstanceInfo.Host)
	if err != nil {
		log.Fatal(err)
	}

	wsUrl := "wss://" + parsedUrl.Host + "/streaming?i=" + c.InstanceInfo.Token

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	socket := gowebsocket.New(wsUrl)

	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Println("Connected to server")
	}

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Println("Received connect error ", err)
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		printNote(message)
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
		return
	}

	socket.Connect()

	uu, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	mainChId := uu.String()

	uu, err = uuid.NewRandom()
	if err != nil {
		return err
	}
	tlChId := uu.String()

	socket.SendText("{\"type\":\"connect\",\"body\":{\"channel\":\"main\",\"id\":\"" + mainChId + "\"}}")

	var channelText string

	if mode == "local" || mode == "global" || mode == "home" {
		channelText = "{\"type\":\"connect\",\"body\":{\"channel\":\"" + mode + "Timeline\",\"id\":\"" + tlChId + "\"}}"
	} else {
		return errors.New("Please select mode in local/home/global")
	}
	socket.SendText(channelText)

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			socket.Close()
			return nil
		}
	}

}

func printNote(message string) {
	var err error

	messageBody, _, _, err := jsonparser.Get([]byte(message), "body", "body")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// とりあえずTextを持ってきてみる
	_, err = jsonparser.GetString(messageBody, "renoteId")

	var note noteData

	if err != nil {
		note, err = pickNote(messageBody)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		_, err = jsonparser.GetString(messageBody, "replyId")
		if err == nil {
			replyParentValue, _, _, _ := jsonparser.Get(messageBody, "reply")
			replyParent, err := pickNote(replyParentValue)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			repStr := fmt.Sprintf("%s \x1b[35m%s(@%s)\x1b[0m\t %s \x1b[32m%s\x1b[0m\x1b[34m(%s)\x1b[0m", replyParent.timestamp, replyParent.name, replyParent.username, replyParent.text, replyParent.attach, replyParent.id)
			fmt.Fprintln(output, repStr)
			note.offset = "    "
		}

	} else { // renoteだったら

		renoteValue, _, _, _ := jsonparser.Get(messageBody, "renote")

		note, err = pickNote(renoteValue)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		note.name = "[RN]" + note.name

	}

	str := fmt.Sprintf("%s%s \x1b[31m%s(@%s)\x1b[0m\t %s \x1b[32m%s\x1b[0m\x1b[34m(%s)\x1b[0m", note.offset, note.timestamp, note.name, note.username, note.text, note.attach, note.id)

	fmt.Fprintln(output, str)

	return
}
