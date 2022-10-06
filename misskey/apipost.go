package misskey

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (c *Client) apiPost(jsonByte []byte, endpoint string) error {

	req, err := http.NewRequest(
		"POST",
		c.InstanceInfo.Host+"/api/"+endpoint,
		bytes.NewBuffer(jsonByte),
	)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	c.resBuf = new(bytes.Buffer)
	if _, err = io.Copy(c.resBuf, resp.Body); err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		fmt.Println(resp.StatusCode, c.resBuf)
		os.Exit(1)
	}
	defer resp.Body.Close()
	return nil
}
